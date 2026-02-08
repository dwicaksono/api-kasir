package app

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"kasir-api/internal/config"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/usecase"

	_ "github.com/lib/pq"
)

func Run() {
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
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	reportRepo := repository.NewReportRepository(db)

	// Context Timeout
	timeout := 10 * time.Second

	// Usecases
	productUsecase := usecase.NewProductUsecase(productRepo, timeout)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo, timeout)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, productRepo, timeout)
	reportUsecase := usecase.NewReportUsecase(reportRepo, timeout)

	// Handlers
	productHandler := handler.NewProductHandler(productUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)
	reportHandler := handler.NewReportHandler(reportUsecase)

	// Router
	mux := http.NewServeMux()

	// Health Check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

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

	// Transaction Routes
	mux.HandleFunc("POST /api/transactions", transactionHandler.CreateTransaction)

	// Report Route
	mux.HandleFunc("GET /api/report", reportHandler.GetReport)

	// Server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}
