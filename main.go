// cmd/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"your-project/internal/handlers"
	"your-project/internal/repository"
	"your-project/internal/service"
	"your-project/pkg/polling"
)

func main() {
	repo := repository.NewProductRepository()
	service := service.NewProductService(repo)
	handler := handlers.NewProductHandler(service)
	productsPoller := polling.NewPoller()

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateProduct(w, r)
			productsPoller.Publish(service.GetProducts())
			fmt.Printf("Total de productos: %d\n", len(service.GetProducts()))
		case http.MethodGet:
			handler.GetProducts(w, r)
		}
	})

	http.HandleFunc("/polling", func(w http.ResponseWriter, r *http.Request) {
		ch := productsPoller.Subscribe()
		defer productsPoller.Unsubscribe(ch)

		select {
		case data := <-ch:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		case <-time.After(30 * time.Second):
			w.WriteHeader(http.StatusNoContent)
		}
	})

	http.HandleFunc("/discount-polling", func(w http.ResponseWriter, r *http.Request) {
		discountProducts := service.GetDiscountProducts()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(discountProducts)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}