package main

import (
	"database/sql"
	"log"
	"net/http"

	"kasir-api/internal/config"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	// Connect to Database
	db, err := sql.Open("postgres", cfg.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	log.Println("Database connected successfully")

	// Dependencies Injection
	// Repositories
	producRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	// Services
	productService := service.NewProductService(producRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	// Handlers
	productHandler := handler.NewProductHandler(productService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Router
	mux := http.NewServeMux()

	// Product Routes
	mux.HandleFunc("GET /api/products", productHandler.GetAll)
	mux.HandleFunc("POST /api/products", productHandler.Create)
	mux.HandleFunc("GET /api/products/{id}", productHandler.GetByID)
	mux.HandleFunc("PUT /api/products/{id}", productHandler.Update)
	mux.HandleFunc("DELETE /api/products/{id}", productHandler.Delete)

	// Category Routes
	mux.HandleFunc("GET /api/categories", categoryHandler.GetAll)
	mux.HandleFunc("POST /api/categories", categoryHandler.Create)
	mux.HandleFunc("GET /api/categories/{id}", categoryHandler.GetByID)
	mux.HandleFunc("PUT /api/categories/{id}", categoryHandler.Update)
	mux.HandleFunc("DELETE /api/categories/{id}", categoryHandler.Delete)

	// Server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}
