package repository

import (
	"context"
	"database/sql"
	"fmt"
	"kasir-api/internal/domain"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO transactions (total_amount, created_at, updated_at)
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id, created_at, updated_at
	`
	err = tx.QueryRowContext(ctx, query, transaction.TotalAmount).Scan(&transaction.ID, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		return err
	}

	for _, detail := range transaction.Details {
		// Decrease stock
		// We use a check on stock >= quantity to ensure we don't go negative
		updateStockQuery := `UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1`
		res, err := tx.ExecContext(ctx, updateStockQuery, detail.Quantity, detail.ProductID)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			// This means either product doesn't exist or stock is insufficient
			return fmt.Errorf("failed to decrease stock for product ID %d: insufficient stock or product not found", detail.ProductID)
		}

		// Insert transaction detail
		detailQuery := `
			INSERT INTO transaction_details (transaction_id, product_id, product_name, quantity, sub_total)
			VALUES ($1, $2, $3, $4, $5)
		`
		_, err = tx.ExecContext(ctx, detailQuery, transaction.ID, detail.ProductID, detail.ProductName, detail.Quantity, detail.SubTotal)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
