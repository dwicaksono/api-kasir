package main

import (
	"context"
	"kasir-api/internal/config"
	"kasir-api/pkg/database"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	dbPool, err := database.Connect(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Could not connect to database for migration: %v", err)
	}
	defer dbPool.Close()

	ctx := context.Background()

	// Create Products Table
	productsQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price BIGINT NOT NULL,
		stock BIGINT NOT NULL
	);`

	_, err = dbPool.Exec(ctx, productsQuery)
	if err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}
	log.Println("products table created or already exists.")

	// Create Categories Table
	categoriesQuery := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT
	);`

	_, err = dbPool.Exec(ctx, categoriesQuery)
	if err != nil {
		log.Fatalf("Failed to create categories table: %v", err)
	}
	log.Println("categories table created or already exists.")

	log.Println("Migration completed successfully.")
}
