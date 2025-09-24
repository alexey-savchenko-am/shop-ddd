package application

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/common"
	"github.com/alexey-savchenko-am/shop-ddd/internal/product/domain"
)

type ChangePriceCommand struct {
	ID    string
	Price int64
}

type ChangePriceCommandHandler struct {
	repo domain.ProductRepository
}

func NewChangePriceCommandHandler(repo domain.ProductRepository) *ChangePriceCommandHandler {
	return &ChangePriceCommandHandler{repo: repo}
}

func (h *ChangePriceCommandHandler) Handle(cmd ChangePriceCommand) common.Result[*domain.Product] {

	id, err := domain.ParseID(cmd.ID)

	if err != nil {
		return common.Failure[*domain.Product](domain.ErrProductInvalidID(cmd.ID))
	}

	findByIdResult := h.repo.ByID(id)

	if !findByIdResult.IsSuccess {
		return common.Failure[*domain.Product](domain.ErrProductNotFound)
	}

	product := findByIdResult.Value

	priceResult := common.NewUsd(cmd.Price)

	if !priceResult.IsSuccess {
		return common.Failure[*domain.Product](*priceResult.Error)
	}

	result := product.ChangePrice(*priceResult.Value)

	if !result.IsSuccess {
		return common.Failure[*domain.Product](*result.Error)
	}

	if err := h.repo.Save(product); err != nil {
		return common.Failure[*domain.Product](common.FromError("product_db_error", err))
	}

	return result
}
