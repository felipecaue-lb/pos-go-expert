package database

import (
	"github.com/felipecaue-lb/goexpert/09-apis/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var error error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		error = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		error = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, error
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	if error := p.DB.First(&product, "id = ?", id).First(&product).Error; error != nil {
		return nil, error
	}

	return &product, nil
}

func (p *Product) Update(product *entity.Product) error {
	if _, error := p.FindByID(product.ID.String()); error != nil {
		return error
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, error := p.FindByID(id)
	if error != nil {
		return error
	}

	return p.DB.Delete(product).Error
}
