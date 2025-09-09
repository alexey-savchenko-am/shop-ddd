package product

import (
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

type Service struct {
	repo product.Repository
}

func NewService(repo product.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProduct(id product.ID, sku, name string, price int) (*product.Product, error) {
	p, err := product.New(id, sku, name, product.Price(price))

	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Service) ChangePrice(id string, newPrice int) error {
	p, err := s.repo.ByID(product.ID(id))
	if err != nil {
		return err
	}

	if err := p.ChangePrice(product.Price(newPrice)); err != nil {
		return err
	}

	return s.repo.Save(p)
}
