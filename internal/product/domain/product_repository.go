package domain

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/common"
)

type ProductRepository interface {
	Save(p *Product) error
	ByID(id ProductID) common.Result[*Product]
}
