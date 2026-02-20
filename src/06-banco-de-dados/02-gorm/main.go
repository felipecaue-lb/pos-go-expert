package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name  string    `gorm:"size:255;not null"`
	Price float64   `gorm:"type:decimal(10,2);not null"`
	gorm.Model
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}

func main() {
	dsn := "root:admin@tcp(mysql-db:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic(error)
	}

	db.AutoMigrate(Product{})

	//create
	/* db.Create(&Product{Name: "Notebook", Price: 3000.00}) */
	/* db.Create(&Product{Name: "Smartwatch", Price: 800.00}) */

	//create bash
	/* products := []Product{
		{Name: "Smartphone", Price: 1500.00},
		{Name: "Tablet", Price: 2000.00},
	}
	db.Create(&products) */

	//find one
	/* var product Product
	//db.First(&product, 1) => Isso funciona se o ID for um número inteiro
	db.First(&product, "name = ?", "Notebook")
	fmt.Println(product) */

	//find all
	/* var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	} */

	/* var products []Product
	db.Limit(2).Offset(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	} */

	// where
	/* var products []Product
	db.Where("price > ?", 1500).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	} */

	/* var products []Product
	db.Where("name LIKE ?", "%smart%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	} */

	//update
	/* var product Product
	db.First(&product, "id = ?", "2dcd2f17-e5b4-40de-b986-71da93450fb2")
	product.Name = "Notebook Gamer"
	db.Save(&product) */

	//delete
	/* var product2 Product
	db.First(&product2, "id = ?", "2dcd2f17-e5b4-40de-b986-71da93450fb2")
	fmt.Println(product2)
	db.Delete(&product2) */
}
