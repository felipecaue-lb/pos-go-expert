package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// create category
	/* category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	category2 := Category{Name: "Informática"}
	db.Create(&category2) */

	// create product
	/* product := Product{Name: "Notebook", Price: 5699.99, Categories: []Category{category, category2}}
	db.Create(&product) */

	var categories []Category
	error = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if error != nil {
		panic(error)
	}
	for _, category := range categories {
		println("Category:", category.Name)
		for _, product := range category.Products {
			println("  Product:", product.Name, "Price:", product.Price)
		}
	}
}
