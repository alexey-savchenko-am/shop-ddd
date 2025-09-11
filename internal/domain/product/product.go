package product

import (
	"fmt"

	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/common"
)

type ID string

type Product struct {
	id    ID
	sku   string
	name  string
	price common.Money
}

func New(id ID, sku, name string, price common.Money) (*Product, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if sku == "" {
		return nil, fmt.Errorf("sku is required")
	}
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	
	return &Product{id: id, sku: sku, name: name, price: price}, nil
}

func (p *Product) ID() ID              { return p.id }
func (p *Product) SKU() string         { return p.sku }
func (p *Product) Name() string        { return p.name }
func (p *Product) Price() common.Money { return p.price }

func (p *Product) ChangePrice(newPrice common.Money) error {
	if newPrice.Currency != p.price.Currency {
		return fmt.Errorf("currency mismatch")
	}
	p.price = newPrice
	return nil
}
