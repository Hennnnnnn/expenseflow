package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Env	string
	Port string	
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config {
		AppName: getEnv("APP_NAME", "ExpenseFlow"),
		Env: getEnv("ENV", "development"),
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}