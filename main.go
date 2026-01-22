package main

import (
	"encoding/json"
	"kasir-api/category"
	"kasir-api/product"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Product Routes
	http.HandleFunc("GET /api/produk", product.GetProduk)
	http.HandleFunc("POST /api/produk", product.PostProduk)
	http.HandleFunc("GET /api/produk/{id}", product.GetProdukById)
	http.HandleFunc("PUT /api/produk/{id}", product.UpdateProduk)
	http.HandleFunc("DELETE /api/produk/{id}", product.DeleteProduk)

	// Category Routes
	http.HandleFunc("GET /api/categories", category.GetCategories)
	http.HandleFunc("POST /api/categories", category.CreateCategory)
	http.HandleFunc("GET /api/categories/{id}", category.GetCategoryById)
	http.HandleFunc("PUT /api/categories/{id}", category.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/{id}", category.DeleteCategory)

	http.ListenAndServe(":8080", nil)
}
