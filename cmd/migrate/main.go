package main

import (
	"fmt"
	"kasir-api/database"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	// Load config manually since this is a standalone script
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Warning: .env file not found")
	}

	connStr := viper.GetString("DB_CONN")
	if connStr == "" {
		// Fallback for when running from root context
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

	fmt.Println("Running migrations...")

	// 1. Rename columns in products table (if they exist)
	// We use IF EXISTS to avoid errors if already migrated
	sqls := []string{
		`ALTER TABLE products RENAME COLUMN nama TO name;`,
		`ALTER TABLE products RENAME COLUMN harga TO price;`,
		`ALTER TABLE products RENAME COLUMN stok TO stock;`,
		// Add description if not exists (handling original schema match)
		`ALTER TABLE products ADD COLUMN IF NOT EXISTS description TEXT;`,
		// Ensure types match what we expect
		`ALTER TABLE products ALTER COLUMN price TYPE DECIMAL(15, 2);`,
		// Create Categories Table
		`CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		// Ensure columns exist (in case table existed but empty/old)
		`ALTER TABLE categories ADD COLUMN IF NOT EXISTS description TEXT;`,
		`ALTER TABLE categories ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;`,
		`ALTER TABLE categories ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;`,

		// Add category_id to products
		`ALTER TABLE products ADD COLUMN IF NOT EXISTS category_id INT;`,

		// Add timestamps to products
		`ALTER TABLE products ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;`,
		`ALTER TABLE products ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;`,

		// Add Foreign Key
		`DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'fk_product_category') THEN 
				ALTER TABLE products ADD CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES categories(id); 
			END IF; 
		END $$;`,
	}

	for _, s := range sqls {
		_, err := db.Exec(s)
		if err != nil {
			// Intentionally log and continue for duplicate rename errors
			// (Postgres doesn't have "RENAME COLUMN IF EXISTS")
			fmt.Printf("Notice/Error executing: %s \nError: %v\n", s, err)
		} else {
			fmt.Println("Executed:", s)
		}
	}

	fmt.Println("Migration finished.")
}
