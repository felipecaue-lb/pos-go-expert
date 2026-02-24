package entity

import (
	"errors"
	"time"

	"github.com/felipecaue-lb/goexpert/09-apis/pkg/entity"
)

var (
	ErrorIDIsRequired    = errors.New("ID is required")
	ErrorInvalidId       = errors.New("invalid ID")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")
	ErrorInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	error := product.Validate()
	if error != nil {
		return nil, error
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequired
	}
	if _, error := entity.ParseID(p.ID.String()); error != nil {
		return ErrorInvalidId
	}

	if p.Name == "" {
		return ErrorNameIsRequired
	}

	if p.Price == 0 {
		return ErrorPriceIsRequired
	}
	if p.Price < 0 {
		return ErrorInvalidPrice
	}

	return nil
}
