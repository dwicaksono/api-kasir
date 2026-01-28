package repository

import (
	"context"
	"database/sql"
	"errors"
	"kasir-api/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	query := `
		SELECT p.id, p.name, p.description, p.price, p.stock, p.category_id, COALESCE(c.name, '') as category_name, p.created_at, p.updated_at 
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
	`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p domain.Product
		var categoryName sql.NullString
		var categoryID sql.NullInt64

		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &categoryID, &categoryName, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		if categoryID.Valid {
			p.CategoryID = int(categoryID.Int64)
		}
		if categoryName.Valid {
			p.CategoryName = categoryName.String
		}
		products = append(products, p)
	}
	return products, nil
}

func (repo *ProductRepository) GetByID(ctx context.Context, id int) (domain.Product, error) {
	var p domain.Product
	var categoryName sql.NullString
	var categoryID sql.NullInt64

	query := `
		SELECT p.id, p.name, p.description, p.price, p.stock, p.category_id, COALESCE(c.name, '') as category_name, p.created_at, p.updated_at 
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1
	`
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &categoryID, &categoryName, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Product{}, errors.New("product not found")
		}
		return domain.Product{}, err
	}
	if categoryID.Valid {
		p.CategoryID = int(categoryID.Int64)
	}
	if categoryName.Valid {
		p.CategoryName = categoryName.String
	}
	return p, nil
}

func (repo *ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	query := `INSERT INTO products (name, description, price, stock, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	err := repo.db.QueryRowContext(ctx, query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) Update(ctx context.Context, id int, product *domain.Product) error {
	query := "UPDATE products SET name = $1, description = $2, price = $3, stock = $4, category_id = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6"
	result, err := repo.db.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (repo *ProductRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("product not found")
	}
	return nil
}
