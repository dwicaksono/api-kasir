package repository

import (
	"context"
	"database/sql"
	"errors"
	"kasir-api/internal/domain"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	query := "SELECT id, name, description, created_at, updated_at FROM categories"
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c domain.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (repo *CategoryRepository) GetByID(ctx context.Context, id int) (domain.Category, error) {
	var c domain.Category
	query := "SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1"
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Category{}, errors.New("category not found")
		}
		return domain.Category{}, err
	}
	return c, nil
}

func (repo *CategoryRepository) Create(ctx context.Context, category *domain.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id, created_at, updated_at"
	err := repo.db.QueryRowContext(ctx, query, category.Name, category.Description).Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CategoryRepository) Update(ctx context.Context, id int, category *domain.Category) error {
	query := "UPDATE categories SET name = $1, description = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3"
	result, err := repo.db.ExecContext(ctx, query, category.Name, category.Description, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (repo *CategoryRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("category not found")
	}
	return nil
}
