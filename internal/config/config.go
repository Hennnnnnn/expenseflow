package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Application
	AppName string
	Env     string
	Port    string

	// IMAP
	IMAPHost     string
	IMAPPort     string
	IMAPUsername string
	IMAPPassword string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		// Application
		AppName: getEnv("APP_NAME", "ExpenseFlow"),
		Env:     getEnv("ENV", "development"),
		Port:    getEnv("PORT", "8080"),

		// IMAP
		IMAPHost:     getEnv("IMAP_HOST", "imap.gmail.com"),
		IMAPPort:     getEnv("IMAP_PORT", "993"),
		IMAPUsername: getEnv("IMAP_USERNAME", ""),
		IMAPPassword: getEnv("IMAP_PASSWORD", ""),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}