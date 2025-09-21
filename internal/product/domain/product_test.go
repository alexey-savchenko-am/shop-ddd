package domain

import (
	"testing"

	"github.com/alexey-savchenko-am/shop-ddd/internal/common"
)

func TestNewProduct(t *testing.T) {

	price, err := common.NewUsd(1000)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = NewProduct("SKU-1", "Product 1", price)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestInvalidProduct(t *testing.T) {

	price, _ := common.NewUsd(0)

	_, err := NewProduct("", "", price)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestChangePrice(t *testing.T) {

	price, _ := common.NewUsd(1000)

	p, _ := NewProduct("SKU-1", "Product 1", price)

	newPrice, _ := common.NewUsd(2000)

	if err := p.ChangePrice(newPrice); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.Price().Amount != newPrice.Amount {
		t.Fatalf("expected price 2000, got %d", p.Price().Amount)
	}
}
