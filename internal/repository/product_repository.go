package repository

import (
	"context"
	"kasir-api/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) domain.ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Fetch(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *productRepository) GetByID(ctx context.Context, id int) (domain.Product, error) {
	var p domain.Product
	err := r.db.QueryRow(ctx, "SELECT id, name, price, stock FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.Product{}, err
		}
		return domain.Product{}, err
	}
	return p, nil
}

func (r *productRepository) Store(ctx context.Context, p *domain.Product) error {
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(ctx, query, p.Name, p.Price, p.Stock).Scan(&p.ID)
	return err
}

func (r *productRepository) Update(ctx context.Context, p *domain.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4"
	_, err := r.db.Exec(ctx, query, p.Name, p.Price, p.Stock, p.ID)
	return err
}

func (r *productRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	return err
}
