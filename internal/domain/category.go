package domain

import "context"

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryRepository interface {
	Fetch(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int) (Category, error)
	Store(ctx context.Context, c *Category) error
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id int) error
}

type CategoryUsecase interface {
	Fetch(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int) (Category, error)
	Store(ctx context.Context, c *Category) error
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id int) error
}
