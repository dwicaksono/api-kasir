package repository

import (
	"context"
	"kasir-api/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) domain.CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Fetch(ctx context.Context) ([]domain.Category, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var c domain.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id int) (domain.Category, error) {
	var c domain.Category
	err := r.db.QueryRow(ctx, "SELECT id, name, description FROM categories WHERE id = $1", id).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.Category{}, err
		}
		return domain.Category{}, err
	}
	return c, nil
}

func (r *categoryRepository) Store(ctx context.Context, c *domain.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(ctx, query, c.Name, c.Description).Scan(&c.ID)
	return err
}

func (r *categoryRepository) Update(ctx context.Context, c *domain.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	_, err := r.db.Exec(ctx, query, c.Name, c.Description, c.ID)
	return err
}

func (r *categoryRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	return err
}
