package product

import (
	"fmt"
	"sync"

	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

type MemoryRepository struct {
	mu sync.RWMutex
	storage map[product.ID]*product.Product
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		storage: make(map[product.ID]*product.Product),
	}
}

func (r *MemoryRepository) Save(p *product.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.storage[p.ID()] = p
	return nil
}

func (r *MemoryRepository) ByID(id product.ID) (*product.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.storage[id]
	if !ok {
		return nil, fmt.Errorf("product %s not found", id)
	}
	return p, nil
}
