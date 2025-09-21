package product

import (
	"fmt"

	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/common"
	"github.com/google/uuid"
)

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	uid, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return ID(uid), nil
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}

type Product struct {
	id    ID
	sku   string
	name  string
	price common.Money
}

func New(sku, name string, price common.Money) (*Product, error) {
	if sku == "" {
		return nil, fmt.Errorf("sku is required")
	}
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if price.Amount <= 0 {
		return nil, fmt.Errorf("price must be > 0")
	}

	return &Product{
		id:    NewID(),
		sku:   sku,
		name:  name,
		price: price,
	}, nil
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
