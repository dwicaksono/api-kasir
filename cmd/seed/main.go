package main

import (
	"database/sql"
	"fmt"
	"kasir-api/database"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	// Load config
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Warning: .env file not found")
	}

	connStr := viper.GetString("DB_CONN")
	if connStr == "" {
		if _, err := os.Stat(".env"); err == nil {
			viper.SetConfigFile(".env")
			viper.ReadInConfig()
			connStr = viper.GetString("DB_CONN")
		}
	}

	if connStr == "" {
		log.Fatal("DB_CONN environment variable is not set")
	}

	db, err := database.ConnectDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Seeding database...")

	// 1. Seed Categories
	categories := []string{
		"Instant Food",
		"Beverages",
		"Snacks",
		"Household",
		"Personal Care",
	}

	var instantFoodID int

	for _, catName := range categories {
		var id int
		err := db.QueryRow("SELECT id FROM categories WHERE name = $1", catName).Scan(&id)
		if err == sql.ErrNoRows {
			err = db.QueryRow("INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id", catName, "Category for "+catName).Scan(&id)
			if err != nil {
				log.Printf("Failed to insert category %s: %v\n", catName, err)
				continue
			}
			fmt.Printf("Inserted category: %s (ID: %d)\n", catName, id)
		} else if err != nil {
			log.Printf("Error checking category %s: %v\n", catName, err)
			continue
		} else {
			fmt.Printf("Category %s already exists (ID: %d)\n", catName, id)
		}

		if catName == "Instant Food" {
			instantFoodID = id
		}
	}

	if instantFoodID == 0 {
		log.Fatal("Failed to get ID for Instant Food category")
	}

	// 2. Seed Indomie Products
	indomieVariants := []string{
		"Indomie Original",
		"Indomie Goreng",
		"Indomie Soto",
		"Indomie Kari Ayam",
		"Indomie Rendang",
		"Indomie Cabe Ijo",
		"Indomie Iga Penyet",
		"Indomie Aceh",
		"Indomie Salted Egg",
		"Indomie Seblak",
	}

	for _, variant := range indomieVariants {
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE name = $1)", variant).Scan(&exists)
		if err != nil {
			log.Printf("Error checking product %s: %v\n", variant, err)
			continue
		}

		if !exists {
			_, err := db.Exec("INSERT INTO products (name, description, price, stock, category_id) VALUES ($1, $2, $3, $4, $5)",
				variant, "Delicious "+variant, 3500, 10, instantFoodID)
			if err != nil {
				log.Printf("Failed to insert product %s: %v\n", variant, err)
			} else {
				fmt.Printf("Inserted product: %s\n", variant)
			}
		} else {
			fmt.Printf("Product %s already exists\n", variant)
		}
	}

	fmt.Println("Seeding completed.")
}
