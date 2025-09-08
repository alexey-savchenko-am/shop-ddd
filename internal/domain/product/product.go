package product

import "fmt"

type ID string
type Price int64

type Product struct {
	id    ID
	sku   string
	name  string
	price Price
}

func New(id ID, sku, name string, price Price) (*Product, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if sku == "" {
		return nil, fmt.Errorf("sku is required")
	}
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if price <= 0 {
		return nil, fmt.Errorf("price must be > 0")
	}
	return &Product{id: id, sku: sku, name: name, price: price}, nil
}

func (p *Product) ID() ID       { return p.id }
func (p *Product) SKU() string  { return p.sku }
func (p *Product) Name() string { return p.name }
func (p *Product) Price() Price { return p.price }

func (p *Product) ChangePrice(newPrice Price) error {
	if(newPrice <= 0) {
		return  fmt.Errorf("new price must be > 0")
	}
	p.price = newPrice
	return nil
}

