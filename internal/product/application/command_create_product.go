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

func (h *CreateProductCommandHandler) Handle(cmd CreateProductCommand) common.Result[*domain.Product] {

	priceResult := common.NewUsd(cmd.Price)

	if !priceResult.IsSuccess {
		return common.Failure[*domain.Product](*priceResult.Error)
	}

	productResult := domain.NewProduct(cmd.SKU, cmd.Name, *priceResult.Value)

	if err := h.repo.Save(productResult.Value); err != nil {
		return common.Failure[*domain.Product](common.FromError("product_db_error", err))
	}

	return productResult
}
