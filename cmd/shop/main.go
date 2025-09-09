package main

import (
	"fmt"

	appProduct "github.com/alexey-savchenko-am/shop-ddd/internal/application/product"
	infraProduct "github.com/alexey-savchenko-am/shop-ddd/internal/infrastructure/product"
)

func main() {
	repo := infraProduct.NewMemoryRepository()
	service := appProduct.NewService(repo)

	// создаём продукт
	p, err := service.CreateProduct("p1", "SKU-1", "Test Product", 1000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created product: %s (%s) — price %d\n", p.Name(), p.SKU(), p.Price())

	// меняем цену
	if err := service.ChangePrice("p1", 2000); err != nil {
		panic(err)
	}

	// читаем обратно из репозитория
	updated, _ := repo.ByID("p1")
	fmt.Printf("Updated product: %s — price %d\n", updated.Name(), updated.Price())
}
