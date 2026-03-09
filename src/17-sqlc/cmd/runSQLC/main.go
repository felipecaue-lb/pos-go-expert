package main

import (
	"context"
	"database/sql"

	"github.com/felipecaue-lb/pos-go-expert/src/17-sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:admin@tcp(mysql-db:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	/* err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:   uuid.New().String(),
		Name: "Backend",
		Description: sql.NullString{
			String: "Backend description",
			Valid:  true,
		},
	})
	if err != nil {
		panic(err)
	} */

	/* err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:   "36ebc4b4-e6b9-4f7a-ab6f-5dbb7e1e6c6e",
		Name: "Backend updated",
		Description: sql.NullString{
			String: "Backend description updated",
			Valid:  true,
		},
	}) */

	err = queries.DeleteCategory(ctx, "7bc8584e-4ad4-438f-a607-d3dcbb932fb7")
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}
