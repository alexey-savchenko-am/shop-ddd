package application

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/common"
	"github.com/alexey-savchenko-am/shop-ddd/internal/product/domain"
)

type CreateProductCommand struct {
	SKU   string
	Name  string
	Price int64
}

type CreateProductCommandHandler struct {
	repo domain.ProductRepository
}

func NewCreateProductCommandHandler(repo domain.ProductRepository) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{repo: repo}
}

func (h *CreateProductCommandHandler) Handle(cmd CreateProductCommand) (*domain.Product, error) {

	price, err := common.NewUsd(cmd.Price)

	if err != nil {
		return nil, err
	}

	newProduct, err := domain.NewProduct(cmd.SKU, cmd.Name, price)

	if err != nil {
		return nil, err
	}

	if err := h.repo.Save(newProduct); err != nil {
		return nil, err
	}

	return newProduct, nil
}
