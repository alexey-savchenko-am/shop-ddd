package main

import (
	"log"
	"net/http"

	_ "github.com/alexey-savchenko-am/shop-ddd/docs" // docs swag init
	httpSwagger "github.com/swaggo/http-swagger"

	appProduct "github.com/alexey-savchenko-am/shop-ddd/internal/application/product"
	"github.com/alexey-savchenko-am/shop-ddd/internal/infrastructure/postgres"
	httpProduct "github.com/alexey-savchenko-am/shop-ddd/internal/interfaces/http/product"
)

func main() {

	db, err := postgres.NewGormDB()

	if err != nil {
		log.Fatal("can not connect db:", err)
	}

	if err := db.AutoMigrate(&postgres.ProductModel{}); err != nil {
		log.Fatal("migration failed", err)
	}

	repo := postgres.NewProductRepository(db)
	useCases := appProduct.NewUseCases(repo)
	handler := httpProduct.NewHandler(useCases)

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateProduct(w, r)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			if id != "" {
				handler.GetById(w, r)
				return
			}
			http.Error(w, "missing id", http.StatusBadRequest)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/products/price", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPatch {
			handler.ChangePrice(w, r)
			return
		}
		http.NotFound(w, r)
	})

	log.Println("Server is running on :3000")
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
