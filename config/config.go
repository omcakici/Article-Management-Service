package config

import (
    "os"
)

// Config holds the configuration for the application
type Config struct {
    ServerPort string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
    return &Config{
        ServerPort: getEnv("SERVER_PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
