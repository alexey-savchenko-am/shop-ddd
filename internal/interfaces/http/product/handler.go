package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	appProduct "github.com/alexey-savchenko-am/shop-ddd/internal/application/product"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	productUseCases *appProduct.UseCases
}

func NewHandler(useCases *appProduct.UseCases) *Handler {
	return &Handler{productUseCases: useCases}
}

type ProductDto struct {
	ID       string `json:"id"`
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Currency string `json:"currency"`
}

type ChangePriceRequest struct {
	NewPrice int64 `json:"new_price"`
}

// CreateProduct godoc
// @Summary      Create product
// @Description  Creates a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body  ProductDto  true  "Product data"
// @Success      201      {object}  ProductDto
// @Failure      400      {string}  string "Bad request"
// @Router       /products [post]
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ProductDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := appProduct.CreateProductCommand{
		ID:    req.ID,
		SKU:   req.SKU,
		Name:  req.Name,
		Price: req.Price,
	}

	created, err := h.productUseCases.CreateProduct.Handle(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := ProductDto{
		ID:       string(created.ID()),
		SKU:      created.SKU(),
		Name:     created.Name(),
		Price:    created.Price().Amount,
		Currency: created.Price().Currency,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

// GetById godoc
// @Summary      Get product by ID
// @Description  Returns product by ID
// @Tags         products
// @Produce      json
// @Param        id   path  string  true  "Product ID"
// @Success      200  {object}  ProductDto
// @Failure      404  {string}  string "Not found"
// @Router       /products/{id} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	
	id := chi.URLParam(r, "id")
	
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	getByIdQuery := appProduct.GetByIdQuery{ID: id}
	p, err := h.productUseCases.GetProductById.Handle(getByIdQuery)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	resp := ProductDto{
		ID:       string(p.ID()),
		SKU:      p.SKU(),
		Name:     p.Name(),
		Price:    p.Price().Amount,
		Currency: p.Price().Currency,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// GetAll godoc
// @Summary Get all products
// @Description Returns list of products with filters and pagination
// @Tags products
// @Produce json
// @Param sku query string false "Filter by SKU"
// @Param name query string false "Filter by name"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ProductDto
// @Router /products [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	var sku *string
	if v := values.Get("sku"); v != "" {
		sku = &v
	}

	var name *string
	if v := values.Get("name"); v != "" {
		name = &v
	}

	limit, err := strconv.Atoi(values.Get("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(values.Get("offset"))
	if err != nil {
		offset = 0
	}

	q := appProduct.GetAllQuery{
		SKU:    sku,
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}

	products, err := h.productUseCases.GetAllProducts.Handle(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(products)
}

// ChangePrice godoc
// @Summary      Change product price
// @Description  Updates the price of a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        path      string  true   "Product ID"
// @Param        request   body      ChangePriceRequest  true  "New price"
// @Success      200       {string}  string "Price updated"
// @Failure      400       {string}  string "Bad Request"
// @Router       /products/{id}/price [patch]
func (h *Handler) ChangePrice(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	var req ChangePriceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	changePriceCmd := appProduct.ChangePriceCommand{
		ID:    id,
		Price: req.NewPrice,
	}

	if _, err := h.productUseCases.ChangePrice.Handle(changePriceCmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("price updated"))
}
