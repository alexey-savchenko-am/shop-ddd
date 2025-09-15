package product

import (
	"fmt"

	"github.com/alexey-savchenko-am/shop-ddd/internal/application/common"
)

type GetAllQuery struct {
	SKU    *string
	Name   *string
	Limit  int
	Offset int
}

type GetAllQueryHandler struct {
	queryDb common.QueryDB
}

func NewGetAllQueryHandler(queryDb common.QueryDB) *GetAllQueryHandler {
	return &GetAllQueryHandler{queryDb: queryDb}
}

type ProductRow struct {
	ID       string `db:"id"`
	SKU      string `db:"sku"`
	Name     string `db:"name"`
	Price    int64  `db:"price"`
	Currency string `db:"currency"`
}

func (h *GetAllQueryHandler) Handle(q GetAllQuery) ([]ProductRow, error) {
	query := `
		SELECT id, sku, name, price, currency
		FROM products
		WHERE 1=1
	`
	args := map[string]any{}

	if q.SKU != nil {
		query += " AND sku = :sku"
		args["sku"] = *q.SKU
	}

	if q.Name != nil {
		query += " AND name ILIKE :name"
		args["name"] = "%" + *q.Name + "%"
	}

	if q.Limit <= 0 {
		q.Limit = 10
	}

	if q.Offset < 0 {
		q.Offset = 0
	}

	query += " ORDER BY name LIMIT :limit OFFSET :offset"

	args["limit"] = q.Limit
	args["offset"] = q.Offset

	
	rows := []ProductRow{}

	err := h.queryDb.Select(&rows, query, args)

	if err != nil {
		return nil, fmt.Errorf("select products: %q", err)
	}

	return rows, nil
}
