package memory

import (
	"fmt"
	"sync"

	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

type ProductRepository struct {
	mu      sync.RWMutex
	storage map[product.ID]*product.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		storage: make(map[product.ID]*product.Product),
	}
}

func (r *ProductRepository) Save(p *product.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.storage[p.ID()] = p
	return nil
}

func (r *ProductRepository) ByID(id product.ID) (*product.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.storage[id]
	if !ok {
		return nil, fmt.Errorf("product %s not found", id)
	}
	return p, nil
}
