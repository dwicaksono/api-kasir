package domain

import (
	"context"
	"time"
)

type Transaction struct {
	ID          int                 `json:"id"`
	TotalAmount float64             `json:"total_amount"`
	Details     []TransactionDetail `json:"details"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type TransactionDetail struct {
	ID            int     `json:"id"`
	TransactionID int     `json:"transaction_id"`
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name,omitempty"`
	Quantity      int     `json:"quantity"`
	SubTotal      float64 `json:"sub_total"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *Transaction) error
}

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, checkoutRequest CheckoutRequest) (*Transaction, error)
}
