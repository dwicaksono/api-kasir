package domain

import "context"

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Stock int64  `json:"stock"`
}

type ProductRepository interface {
	Fetch(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id int) (Product, error)
	Store(ctx context.Context, p *Product) error
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id int) error
}

type ProductUsecase interface {
	Fetch(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id int) (Product, error)
	Store(ctx context.Context, p *Product) error
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id int) error
}
