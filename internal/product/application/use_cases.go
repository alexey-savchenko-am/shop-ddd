package application

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/common/persistence"
	"github.com/alexey-savchenko-am/shop-ddd/internal/product/domain"
)

type UseCases struct {
	CreateProduct  *CreateProductCommandHandler
	ChangePrice    *ChangePriceCommandHandler
	GetProductById *GetByIdQueryHandler
	GetAllProducts *GetAllQueryHandler
}

func NewUseCases(queryDb persistence.QueryDB, repo domain.ProductRepository) *UseCases {
	return &UseCases{
		CreateProduct:  NewCreateProductCommandHandler(repo),
		ChangePrice:    NewChangePriceCommandHandler(repo),
		GetProductById: NewGetByIdQueryHandler(queryDb),
		GetAllProducts: NewGetAllQueryHandler(queryDb),
	}
}
