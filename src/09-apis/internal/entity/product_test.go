package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, error := NewProduct("Product 1", 10.00)
	assert.Nil(t, error)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Price)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.00, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, error := NewProduct("", 10.00)
	assert.Nil(t, product)
	assert.Equal(t, ErrorNameIsRequired, error)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, error := NewProduct("Product 1", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrorPriceIsRequired, error)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, error := NewProduct("Product 1", -1)
	assert.Nil(t, product)
	assert.Equal(t, ErrorInvalidPrice, error)
}

func TestProductValidate(t *testing.T) {
	product, error := NewProduct("Product 1", 10.00)
	assert.Nil(t, error)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
