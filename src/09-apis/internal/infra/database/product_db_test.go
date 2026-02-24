package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/felipecaue-lb/goexpert/09-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, error := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if error != nil {
		t.Error(error)
	}

	db.AutoMigrate(&entity.Product{})

	product, error := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, error)

	productDB := NewProduct(db)
	error = productDB.Create(product)
	assert.NoError(t, error)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, error := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if error != nil {
		t.Error(error)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i <= 24; i++ {
		product, error := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, error)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, error := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, error)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, error = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, error)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, error = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, error)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)
}

func TestFindByIDProduct(t *testing.T) {
	db, error := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if error != nil {
		t.Error(error)
	}

	db.AutoMigrate(&entity.Product{})

	product, error := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, error)
	db.Create(product)
	productDB := NewProduct(db)

	foundProduct, error := productDB.FindByID(product.ID.String())
	assert.NoError(t, error)
	assert.Equal(t, product.ID, foundProduct.ID)
	assert.Equal(t, product.Name, foundProduct.Name)
	assert.Equal(t, product.Price, foundProduct.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, error := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if error != nil {
		t.Error(error)
	}

	db.AutoMigrate(&entity.Product{})

	product, error := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, error)
	db.Create(product)
	productDB := NewProduct(db)

	product.Name = "Updated Product"
	product.Price = 20.00
	error = productDB.Update(product)
	assert.NoError(t, error)

	updatedProduct, error := productDB.FindByID(product.ID.String())
	assert.NoError(t, error)
	assert.Equal(t, "Updated Product", updatedProduct.Name)
	assert.Equal(t, 20.00, updatedProduct.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, error := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if error != nil {
		t.Error(error)
	}

	db.AutoMigrate(&entity.Product{})

	product, error := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, error)
	db.Create(product)
	productDB := NewProduct(db)

	error = productDB.Delete(product.ID.String())
	assert.NoError(t, error)

	deletedProduct, error := productDB.FindByID(product.ID.String())
	assert.Error(t, error)
	assert.Nil(t, deletedProduct)
}
