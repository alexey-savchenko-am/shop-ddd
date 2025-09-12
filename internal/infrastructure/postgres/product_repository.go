package postgres

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/common"
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(p *product.Product) error {
	model := ProductModel{
		ID:       string(p.ID()),
		SKU:      p.SKU(),
		Name:     p.Name(),
		Price:    p.Price().Amount,
		Currency: p.Price().Currency,
	}

	return r.db.Save(&model).Error
}

func (r *ProductRepository) ByID(id product.ID) (*product.Product, error) {
	var model ProductModel

	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}

	money := common.Money{Amount: model.Price, Currency: model.Currency}
	return product.New(product.ID(model.ID), model.SKU, model.Name, money)
}
