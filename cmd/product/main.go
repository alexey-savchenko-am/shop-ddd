package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/alexey-savchenko-am/shop-ddd/docs" // docs swag init
	httpSwagger "github.com/swaggo/http-swagger"

	pers  "github.com/alexey-savchenko-am/shop-ddd/internal/common/persistence"
	prodApp "github.com/alexey-savchenko-am/shop-ddd/internal/product/application"
	prodPersistence "github.com/alexey-savchenko-am/shop-ddd/internal/product/infrastructure/persistence"
	httpHandlers "github.com/alexey-savchenko-am/shop-ddd/internal/product/interfaces/http"
)

func main() {

	db, err := pers.NewGormDB()

	if err != nil {
		log.Fatal("can not connect db:", err)
	}

	sqlxDB, err := pers.NewSqlxDB()

	if err != nil {
		log.Fatal("can not connect db:", err)
	}

	queryDb := pers.NewSqlxQueryDB(sqlxDB)
	repo := prodPersistence.NewProductRepository(db)
	useCases := prodApp.NewUseCases(queryDb, repo)
	handler := httpHandlers.NewHandler(useCases)

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
