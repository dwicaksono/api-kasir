package domain

import (
	"context"
	"time"
)

type Product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	CategoryID   int       `json:"category_id"`
	CategoryName string    `json:"category_name,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ProductRepository interface {
	GetAll(ctx context.Context, name string) ([]Product, error)
	GetByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, id int, product *Product) error
	Delete(ctx context.Context, id int) error
}

type ProductUsecase interface {
	GetAll(ctx context.Context, name string) ([]Product, error)
	GetByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, id int, product *Product) error
	Delete(ctx context.Context, id int) error
}
