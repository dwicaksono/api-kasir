package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBConnection string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Warning: .env file not found or cannot be read, relying on system environment variables")
	}

	config := Config{
		Port:         viper.GetString("PORT"),
		DBConnection: viper.GetString("DB_CONN"),
	}

	db, err := database.ConnectDB(config.DBConnection)
	if err != nil {
		log.Fatal("gagal connect database ", err)
	}
	defer db.Close()

	// Health Check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Initialize Product Components
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Initialize Category Components
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Product Routes
	http.HandleFunc("GET /api/products", productHandler.GetAll)
	http.HandleFunc("POST /api/products", productHandler.Create)
	http.HandleFunc("GET /api/products/{id}", productHandler.GetByID)
	http.HandleFunc("PUT /api/products/{id}", productHandler.Update)
	http.HandleFunc("DELETE /api/products/{id}", productHandler.Delete)

	// Category Routes
	http.HandleFunc("GET /api/categories", categoryHandler.GetAll)
	http.HandleFunc("POST /api/categories", categoryHandler.Create)
	http.HandleFunc("GET /api/categories/{id}", categoryHandler.GetByID)
	http.HandleFunc("PUT /api/categories/{id}", categoryHandler.Update)
	http.HandleFunc("DELETE /api/categories/{id}", categoryHandler.Delete)

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
