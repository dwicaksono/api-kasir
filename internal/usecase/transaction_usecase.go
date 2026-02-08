package usecase

import (
	"context"
	"errors"
	"kasir-api/internal/domain"
	"time"
)

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
	productRepo     domain.ProductRepository
	contextTimeout  time.Duration
}

func NewTransactionUsecase(tr domain.TransactionRepository, pr domain.ProductRepository, timeout time.Duration) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: tr,
		productRepo:     pr,
		contextTimeout:  timeout,
	}
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, checkoutRequest domain.CheckoutRequest) (*domain.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var totalAmount float64
	var details []domain.TransactionDetail

	for _, item := range checkoutRequest.Items {
		product, err := u.productRepo.GetByID(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		subTotal := product.Price * float64(item.Quantity)
		totalAmount += subTotal

		details = append(details, domain.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			SubTotal:    subTotal,
		})
	}

	transaction := &domain.Transaction{
		TotalAmount: totalAmount,
		Details:     details,
	}

	err := u.transactionRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	// Ideally, we should also save transaction details and update stock here,
	// clearly wrapped in a database transaction.
	// For now, based on the scope, we proceed with creating the transaction header.

	return transaction, nil
}
