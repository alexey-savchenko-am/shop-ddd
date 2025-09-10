package product

import (
	"encoding/json"
	"net/http"

	appProduct "github.com/alexey-savchenko-am/shop-ddd/internal/application/product"
	"github.com/alexey-savchenko-am/shop-ddd/internal/domain/product"
)

type Handler struct {
	service *appProduct.Service
}

type ProductDto struct {
	ID    product.ID `json:"id"`
	SKU   string     `json:"sku"`
	Name  string     `json:"name"`
	Price int        `json:"price"`
}

func NewHandler(service *appProduct.Service) *Handler {
	return &Handler{service: service}
}

// CreateProduct godoc
// @Summary      Create product
// @Description  Creates a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body  ProductDto  true  "Product data"
// @Success      200      {object}  ProductDto
// @Failure      400      {string}  string "Bad request"
// @Router       /products [post]
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ProductDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := h.service.CreateProduct(req.ID, req.SKU, req.Name, req.Price)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", "/products/"+string(p.ID()))
	_ = json.NewEncoder(w).Encode(map[string]any{
		"id":    p.ID(),
		"sku":   p.SKU(),
		"name":  p.Name(),
		"price": p.Price(),
	})
}

// GetById godoc
// @Summary      Get product by ID
// @Description  Returns product by ID
// @Tags         products
// @Produce      json
// @Param        id   query  string  true  "Product ID"
// @Success      200  {object}  ProductDto
// @Failure      404  {string}  string "Not found"
// @Router       /products [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	productId := product.ID(id)

	p, err := h.service.GetById(productId)
	
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	resp := ProductDto{
		ID:    p.ID(),
		SKU:   p.SKU(),
		Name:  p.Name(),
		Price: int(p.Price()),
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// PATCH /products/{id}/price
func (h *Handler) ChangePrice(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	var req struct {
		NewPrice int `json:"new_price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.ChangePrice(id, req.NewPrice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
