package product

import (
	"fmt"

	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/common"
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

type ChangePriceCommand struct {
	ID    string
	Price int64
}

type ChangePriceCommandHandler struct {
	repo product.Repository
}

func NewChangePriceCommandHandler(repo product.Repository) *ChangePriceCommandHandler {
	return &ChangePriceCommandHandler{repo: repo}
}

func (h *ChangePriceCommandHandler) Handle(cmd ChangePriceCommand) (*product.Product, error) {

	id, err := product.ParseID(cmd.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid product id: %w", err)
	}
	
	p, err := h.repo.ByID(id)

	if err != nil {
		return nil, err
	}

	newPrice, err := common.NewUsd(cmd.Price)

	if err != nil {
		return nil, err
	}

	if err := p.ChangePrice(newPrice); err != nil {
		return nil, err
	}

	return p, h.repo.Save(p)
}
