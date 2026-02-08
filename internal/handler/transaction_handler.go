package handler

import (
	"encoding/json"
	"kasir-api/internal/domain"
	"net/http"
)

type TransactionHandler struct {
	usecase domain.TransactionUsecase
}

func NewTransactionHandler(usecase domain.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{usecase: usecase}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var checkoutRequest domain.CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&checkoutRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transaction, err := h.usecase.CreateTransaction(r.Context(), checkoutRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) CheckoutTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var checkoutRequest domain.CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&checkoutRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transaction, err := h.usecase.CreateTransaction(r.Context(), checkoutRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(transaction)
}