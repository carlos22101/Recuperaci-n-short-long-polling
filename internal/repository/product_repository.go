// internal/repository/product_repository.go
package repository

import (
	"sync"
	"your-project/internal/domain"
)

type ProductRepository struct {
	mu       sync.RWMutex
	products []domain.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []domain.Product{},
	}
}

func (r *ProductRepository) Add(product domain.Product) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products = append(r.products, product)
}

func (r *ProductRepository) GetAll() []domain.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]domain.Product{}, r.products...)
}

func (r *ProductRepository) GetDiscountProducts() []domain.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	discountProducts := []domain.Product{}
	for _, p := range r.products {
		if p.Discount {
			discountProducts = append(discountProducts, p)
		}
	}
	return discountProducts
}