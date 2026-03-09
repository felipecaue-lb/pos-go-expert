package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/felipecaue-lb/pos-go-expert/src/17-sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:admin@tcp(mysql-db:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Println("-----------------------------------------------")
		fmt.Printf("Category: %s\n", course.CategoryName)
		fmt.Printf("Course ID: %s\n", course.ID)
		fmt.Printf("Course Name: %s\n", course.Name)
		fmt.Printf("Course Description: %s\n", course.Description.String)
		fmt.Printf("Course Price: %f\n", course.Price)
	}

	/* courseArgs := CourseParams{
		ID:   uuid.New().String(),
		Name: "GO",
		Description: sql.NullString{
			String: "GO Course",
			Valid:  true,
		},
		Price: 100.50,
	}

	categoryArgs := CategoryParams{
		ID:   uuid.New().String(),
		Name: "Backend",
		Description: sql.NullString{
			String: "Backend Course",
			Valid:  true,
		},
	}

	courseDB := NewCourseDB(dbConn)
	err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	if err != nil {
		panic(err)
	} */
}
