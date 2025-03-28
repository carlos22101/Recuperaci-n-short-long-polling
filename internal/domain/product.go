
package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Code      string    `json:"code"`
	Discount  bool      `json:"discount"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int, code string, discount bool) *Product {
	return &Product{
		ID:        uuid.New(),
		Name:      name,
		Price:     price,
		Code:      code,
		Discount:  discount,
		CreatedAt: time.Now(),
	}
}