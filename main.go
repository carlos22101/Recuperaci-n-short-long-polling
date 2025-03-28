package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"recu/internal/handlers"
	"recu/internal/repository"
	"recu/internal/service"
	"recu/polling"
)

func main() {
	// Configurar logging
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Crear componentes
	repo := repository.NewProductRepository()
	service := service.NewProductService(repo)
	handler := handlers.NewProductHandler(service)
	productsPoller := polling.NewPoller()

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		// Configurar CORS para pruebas locales
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		switch r.Method {
		case http.MethodPost:
			log.Println("Recibiendo nuevo producto...")
			handler.CreateProduct(w, r)
			productos := service.GetProducts()
			productsPoller.Publish(productos)
			log.Printf(" Total de productos: %d\n", len(productos))
		case http.MethodGet:
			log.Println(" Solicitando lista de productos")
			handler.GetProducts(w, r)
		}
	})

	http.HandleFunc("/polling", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		ch := productsPoller.Subscribe()
		defer func() {
			productsPoller.Unsubscribe(ch)
			close(ch)
		}()

		log.Println("Iniciando polling de productos")
		select {
		case data := <-ch:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			log.Println(" Datos enviados en polling")
		case <-time.After(30 * time.Second):
			w.WriteHeader(http.StatusNoContent)
			log.Println("Timeout en polling")
		}
	})

	http.HandleFunc("/discount-polling", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		discountProducts := service.GetDiscountProducts()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(discountProducts)
		
		log.Printf("Productos con descuento: %d\n", len(discountProducts))
	})


	const PORT = ":8080"
	log.Printf("Servidor corriendo en http://localhost%s\n", PORT)


	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}