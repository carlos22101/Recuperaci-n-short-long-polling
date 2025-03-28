// internal/service/product_service.go
package service

import (
	"recu/internal/domain"
	"recu/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(name string, price int, code string, discount bool) *domain.Product {
	product := domain.NewProduct(name, price, code, discount)
	s.repo.Add(*product)
	return product
}

func (s *ProductService) GetProducts() []domain.Product {
	return s.repo.GetAll()
}

func (s *ProductService) GetDiscountProducts() []domain.Product {
	return s.repo.GetDiscountProducts()
}