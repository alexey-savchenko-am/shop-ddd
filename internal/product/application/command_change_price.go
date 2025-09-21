package application

import (
	"fmt"

	"github.com/alexey-savchenko-am/shop-ddd/internal/product/domain"
	"github.com/alexey-savchenko-am/shop-ddd/internal/common"
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

func (h *ChangePriceCommandHandler) Handle(cmd ChangePriceCommand) (*domain.Product, error) {

	id, err := domain.ParseID(cmd.ID)
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
