package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int       `gorm:"primaryKey;autoIncrement"`
	Name     string    `gorm:"size:255;not null"`
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int        `gorm:"primaryKey;autoIncrement"`
	Name       string     `gorm:"size:255;not null"`
	Price      float64    `gorm:"type:decimal(10,2);not null"`
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:admin@tcp(mysql-db:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic(error)
	}

	db.AutoMigrate(&Product{}, &Category{})

	tx := db.Begin()
	var category Category
	error = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 1).Error
	if error != nil {
		panic(error)
	}

	category.Name = "Eletrônicos e Informática"
	tx.Debug().Save(&category)

	tx.Commit()
}
