package domain

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/common"
	"github.com/google/uuid"
)

type ProductID uuid.UUID

func NewProductID() ProductID {
	return ProductID(uuid.New())
}

func ParseID(s string) (ProductID, error) {
	uid, err := uuid.Parse(s)
	if err != nil {
		return ProductID{}, err
	}
	return ProductID(uid), nil
}

func (id ProductID) String() string {
	return uuid.UUID(id).String()
}

type Product struct {
	id    ProductID
	sku   string
	name  string
	price common.Money
}

func ReconstituteProduct(id ProductID, sku, name string, price common.Money) common.Result[*Product] {

	if sku == "" {
		return common.Failure[*Product](ErrProductSKURequired)
	}
	if name == "" {
		return common.Failure[*Product](ErrProductNameRequired)
	}
	if price.Amount <= 0 {
		return common.Failure[*Product](ErrProductInvalidPrice)
	}

	product := &Product{
		id:    id,
		sku:   sku,
		name:  name,
		price: price,
	}

	return common.Success[*Product](product)
}

func NewProduct(sku, name string, price common.Money) common.Result[*Product] {
	return ReconstituteProduct(NewProductID(), sku, name, price)
}

func (p *Product) ID() ProductID       { return p.id }
func (p *Product) SKU() string         { return p.sku }
func (p *Product) Name() string        { return p.name }
func (p *Product) Price() common.Money { return p.price }

func (p *Product) ChangePrice(newPrice common.Money) common.Result[*Product] {
	if newPrice.Currency != p.price.Currency {
		return common.Failure[*Product](ErrProductInvalidCurrency)
	}
	p.price = newPrice
	return common.Success[*Product](p)
}
