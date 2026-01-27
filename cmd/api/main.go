package main

import (
	"kasir-api/internal/config"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/usecase"
	"kasir-api/pkg/database"
	"log"
	"net/http"
	"time"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	// Initialize Database
	dbPool, err := database.Connect(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer dbPool.Close()

	// Initialize Repositories
	productRepo := repository.NewProductRepository(dbPool)
	categoryRepo := repository.NewCategoryRepository(dbPool)

	// Initialize Usecases
	timeout := 2 * time.Second
	productUsecase := usecase.NewProductUsecase(productRepo, timeout)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo, timeout)

	// Initialize Router
	mux := http.NewServeMux()

	// Health Check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"OK", "message":"API Running"}`))
	})

	// Initialize Handlers
	handler.NewProductHandler(mux, productUsecase)
	handler.NewCategoryHandler(mux, categoryUsecase)

	// Start Server
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
