package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBUrl         string `mapstructure:"DB_CONN"`
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	if config.ServerAddress == "" {
		port := viper.GetString("PORT")
		if port != "" {
			config.ServerAddress = ":" + port
		} else {
			config.ServerAddress = ":8080"
		}
	}

	return &config
}
