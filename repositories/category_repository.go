package repositories

import (
	"database/sql"
	"kasir-api/models"
	"errors"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	rows, err := repo.db.Query("SELECT id, name, description, created_at, updated_at FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (repo *CategoryRepository) GetByID(id int) (models.Category, error) {
	var c models.Category
	err := repo.db.QueryRow("SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1", id).Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return models.Category{}, err
	}
	return c, nil
}

func (repo *CategoryRepository) Create(category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id, name, description, created_at, updated_at"
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CategoryRepository) Update(id int, category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2, updated_at = $3 WHERE id = $4"
	result, err := repo.db.Exec(query, category.Name, category.Description, category.UpdatedAt, id)
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

func (repo *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.Exec(query, id)
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