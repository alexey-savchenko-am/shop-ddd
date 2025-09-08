package product

import "testing"

func TestNewProduct(t *testing.T) {
	_, err := New("p1", "SKU-1", "Product 1", Price(1000))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestInvalidProduct(t *testing.T) {
	_, err := New("", "", "", 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestChangePrice(t *testing.T) {
	p, _ := New("p1", "SKU-1", "Product 1", Price(1000))

	if err := p.ChangePrice(2000); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Price() != 2000 {
		t.Fatalf("expected price 2000, got %d", p.Price())
	}

	if err := p.ChangePrice(0); err == nil {
		t.Fatalf("expected error, got nil")
	}
}

