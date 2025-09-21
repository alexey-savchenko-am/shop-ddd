package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/alexey-savchenko-am/shop-ddd/docs" // docs swag init
	httpSwagger "github.com/swaggo/http-swagger"

	appProduct "github.com/alexey-savchenko-am/shop-ddd/internal/application/product"
	"github.com/alexey-savchenko-am/shop-ddd/internal/infrastructure/persistence"
	"github.com/alexey-savchenko-am/shop-ddd/internal/infrastructure/postgres"
	httpProduct "github.com/alexey-savchenko-am/shop-ddd/internal/interfaces/http/product"
)

func main() {

	db, err := postgres.NewGormDB()

	if err != nil {
		log.Fatal("can not connect db:", err)
	}

	sqlxDB, err := postgres.NewSqlxDB()

	if err != nil {
		log.Fatal("can not connect db:", err)
	}

	queryDb := persistence.NewSqlxQueryDB(sqlxDB)
	repo := postgres.NewProductRepository(db)
	useCases := appProduct.NewUseCases(queryDb, repo)
	handler := httpProduct.NewHandler(useCases)

	// chi router

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// routes

	r.Post("/products", handler.CreateProduct)
	r.Get("/products", handler.GetAll)
	r.Get("/products/{id}", handler.GetById)
	r.Patch("/products/{id}/price", handler.ChangePrice)

	//swagger
	r.Mount("/swagger", httpSwagger.WrapHandler)

	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
