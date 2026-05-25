package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	AdminAPIKey string
	GinMode     string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://booking:booking@localhost:5432/booking?sslmode=disable"),
		AdminAPIKey: getEnv("ADMIN_API_KEY", "admin-secret-key"),
		GinMode:     getEnv("GIN_MODE", "debug"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
