package main

import (
	"fmt"

	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

func main() {
	p, err := product.New("p1", "SKU-1", "Test Product", 1000)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New product created: %s (%s) — price %d\n", p.Name(), p.SKU(), p.Price())
}
