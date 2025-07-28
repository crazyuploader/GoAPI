package config

import (
	"os"
)

type Config struct {
	AppName     string
	Port        string
	Environment string
}

func Load() *Config {
	return &Config{
		AppName:     getEnv("APP_NAME", "My API"),
		Port:        getEnv("PORT", "3100"),
		Environment: getEnv("ENVIRONMENT", "Development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
