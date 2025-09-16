package product

import (
	"fmt"

	"github.com/alexey-savchenko-am/shop-ddd/internal/application/common"
)

type GetByIdQuery struct {
	ID string
}

type GetByIdQueryHandler struct {
	queryDb common.QueryDB
}

func NewGetByIdQueryHandler(queryDb common.QueryDB) *GetByIdQueryHandler {
	return &GetByIdQueryHandler{queryDb: queryDb}
}

func (h *GetByIdQueryHandler) Handle(q GetByIdQuery) (*ProductRow, error) {

	query := `
		SELECT id, sku, name, price, currency
		FROM products
		WHERE id = :id
	`
	args := map[string]any{
		"id": q.ID,
	}

	var row ProductRow

	if err := h.queryDb.Get(&row, query, args); err != nil {
		return nil, fmt.Errorf("get product by id: %w", err)
	}

	return &row, nil
}
