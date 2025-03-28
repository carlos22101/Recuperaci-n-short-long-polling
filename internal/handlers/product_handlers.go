// internal/handlers/product_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"your-project/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name      string `json:"name"`
		Price     int    `json:"price"`
		Code      string `json:"code"`
		Discount  bool   `json:"discount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := h.service.CreateProduct(req.Name, req.Price, req.Code, req.Discount)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetProducts()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}