package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	DBConn string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	viper.AutomaticEnv()

	port := viper.GetString("PORT")
	if port == "" {
		port = ":8080"
	}
	// Ensure port starts with colon
	if port[0] != ':' {
		port = ":" + port
	}

	dbConn := viper.GetString("DB_CONN")
	if dbConn == "" {
		log.Fatal("DB_CONN environment variable is required")
	}

	return &Config{
		Port:   port,
		DBConn: dbConn,
	}
}
