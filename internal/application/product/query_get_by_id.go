package product

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

type GetByIdQuery struct {
	ID string
}

type GetByIdQueryHandler struct {
	repo product.Repository
}

func NewGetByIdQueryHandler(repo product.Repository) *GetByIdQueryHandler {
	return &GetByIdQueryHandler{repo: repo}
}

func (h *GetByIdQueryHandler) Handle(q GetByIdQuery) (*product.Product, error) {

	p, err := h.repo.ByID(product.ID(q.ID))

	if err != nil {
		return nil, err
	}

	return p, nil
}
