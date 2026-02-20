package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int       `gorm:"primaryKey;autoIncrement"`
	Name     string    `gorm:"size:255;not null"`
	Products []Product // has many relationship
}

type Product struct {
	ID           int          `gorm:"primaryKey;autoIncrement"`
	Name         string       `gorm:"size:255;not null"`
	Price        float64      `gorm:"type:decimal(10,2);not null"`
	CategoryID   int          `gorm:"not null"`
	Category     Category     // belongs to relationship
	SerialNumber SerialNumber // has one relationship
	gorm.Model
}

type SerialNumber struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Number    string `gorm:"size:255;not null"`
	ProductID int    `gorm:"not null"`
}

func main() {
	dsn := "root:admin@tcp(mysql-db:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic(error)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	/* category := Category{Name: "Cozinha"}
	db.Create(&category) */

	// create product
	/* product := Product{Name: "Panela", Price: 99.99, CategoryID: 1}
	db.Create(&product) */

	// create serial number
	/* serialNumber := SerialNumber{Number: "SN123456789", ProductID: 1}
	db.Create(&serialNumber) */

	/* var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println("Product:", product.Name, "Price:", product.Price, "Category:", product.Category.Name, "Serial Number:", product.SerialNumber.Number)
	} */

	var categories []Category
	error = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if error != nil {
		panic(error)
	}
	for _, category := range categories {
		println("Category:", category.Name)
		for _, product := range category.Products {
			println("  Product:", product.Name, "Price:", product.Price, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
