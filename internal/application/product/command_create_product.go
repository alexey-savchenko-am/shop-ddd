package product

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/common"
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

type CreateProductCommand struct {
	SKU   string
	Name  string
	Price int64
}

type CreateProductCommandHandler struct {
	repo product.Repository
}

func NewCreateProductCommandHandler(repo product.Repository) *CreateProductCommandHandler {
    return &CreateProductCommandHandler{repo: repo}
}

func (h *CreateProductCommandHandler) Handle(cmd CreateProductCommand) (*product.Product, error) {

	price, err := common.NewUsd(cmd.Price)

	if err != nil {
		return nil, err
	}

	newProduct, err := product.New(cmd.SKU, cmd.Name, price)

	if err != nil {
		return nil, err
	}

	if err := h.repo.Save(newProduct); err != nil {
		return nil, err
	}

	return newProduct, nil
}
