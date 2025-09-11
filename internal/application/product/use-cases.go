package product

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

type UseCases struct {
	CreateProduct *CreateProductCommandHandler
	ChangePrice *ChangePriceCommandHandler
	GetProductById *GetByIdQueryHandler
}

func NewUseCases(repo product.Repository) *UseCases {
	return &UseCases{
		CreateProduct: NewCreateProductCommandHandler(repo),
		ChangePrice: NewChangePriceCommandHandler(repo),
		GetProductById: NewGetByIdQueryHandler(repo),
	}
}