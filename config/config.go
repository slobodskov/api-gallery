package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration parameters for the application
type Config struct {
	DatabasePath string
	ServerPort   string
	MaxFileSize  int64
	GinMode      string
}

// Load initializes and returns the application configuration
// It loads values from environment variables with fallback to defaults
func Load() (*Config, error) {
	godotenv.Load()

	return &Config{
		DatabasePath: getEnv("DB_PATH", "./gallery.db"),
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		MaxFileSize:  getEnvAsInt64("MAX_FILE_SIZE", 10*1024*1024),
		GinMode:      getEnv("GIN_MODE", "debug"),
	}, nil
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt64 retrieves environment variable as int64 or returns default value
func getEnvAsInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}
