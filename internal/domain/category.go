package domain

import (
	"context"
	"time"
)

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int) (Category, error)
	Create(ctx context.Context, category *Category) error
	Update(ctx context.Context, id int, category *Category) error
	Delete(ctx context.Context, id int) error
}

type CategoryUsecase interface {
	GetAll(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int) (Category, error)
	Create(ctx context.Context, category *Category) error
	Update(ctx context.Context, id int, category *Category) error
	Delete(ctx context.Context, id int) error
}
