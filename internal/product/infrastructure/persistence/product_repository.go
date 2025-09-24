package postgres

import (
	"gorm.io/gorm"

	"github.com/alexey-savchenko-am/shop-ddd/internal/common"
	"github.com/alexey-savchenko-am/shop-ddd/internal/product/domain"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(p *domain.Product) error {
	model := ProductModel{
		ID:       p.ID().String(),
		SKU:      p.SKU(),
		Name:     p.Name(),
		Price:    p.Price().Amount,
		Currency: p.Price().Currency,
	}

	return r.db.Save(&model).Error
}

func (r *ProductRepository) ByID(id domain.ProductID) common.Result[*domain.Product] {
	var model ProductModel

	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return common.Failure[*domain.Product](domain.ErrProductNotFound)
	}

	money := common.Money{Amount: model.Price, Currency: model.Currency}

	productId, _ := domain.ParseID(model.ID)

	return domain.ReconstituteProduct(productId, model.SKU, model.Name, money)
}
