/*
* Rodar esses comandos para instalar a dependência do uuid
* go mod init github.com/felipecaue-lb/goexpert/6/1
* go mod tidy
 */

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, error := sql.Open("mysql", "root:admin@tcp(mysql-db:3306)/goexpert")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	product := NewProduct("Notebook", 1899.90)
	error = insertProduct(db, product)
	if error != nil {
		panic(error)
	}

	product.Price = 2000.90
	error = updateProduct(db, product)
	if error != nil {
		panic(error)
	}

	findOneProduct, error := findOneProduct(db, product.ID)
	if error != nil {
		panic(error)
	}
	fmt.Printf("Produto: %v, possui o preço de %.2f\n", findOneProduct.Name, findOneProduct.Price)

	products, error := findAllProducts(db)
	if error != nil {
		panic(error)
	}
	for _, product := range products {
		fmt.Printf("Produto: %v, possui o preço de %.2f\n", product.Name, product.Price)
	}

	error = deleteProduct(db, product.ID)
	if error != nil {
		panic(error)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	smtp, error := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if error != nil {
		return error
	}
	defer smtp.Close()

	_, error = smtp.Exec(product.ID, product.Name, product.Price)
	if error != nil {
		return error
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	smtp, error := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer smtp.Close()

	_, error = smtp.Exec(product.Name, product.Price, product.ID)
	if error != nil {
		return error
	}

	return nil
}

func findOneProduct(db *sql.DB, id string) (*Product, error) {
	smtp, error := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if error != nil {
		return nil, error
	}
	defer smtp.Close()

	var product Product
	error = smtp.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if error != nil {
		return nil, error
	}

	return &product, nil
}

func findAllProducts(db *sql.DB) ([]Product, error) {
	rows, error := db.Query("SELECT id, name, price FROM products")
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		error = rows.Scan(&product.ID, &product.Name, &product.Price)
		if error != nil {
			return nil, error
		}
		products = append(products, product)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	smtp, error := db.Prepare("DELETE FROM products WHERE id = ?")
	if error != nil {
		return error
	}
	defer smtp.Close()

	_, error = smtp.Exec(id)
	if error != nil {
		return error
	}

	return nil
}
