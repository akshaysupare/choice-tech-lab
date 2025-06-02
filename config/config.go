package config

import (
	"os"
)

// Config holds all configuration values for the application
// including MySQL and Redis connection details.
type Config struct {
	MySQLDSN      string // MySQL Data Source Name (user:pass@tcp(host:port)/dbname)
	RedisAddr     string // Redis server address (host:port)
	RedisPassword string // Redis password (if any)
	RedisDB       int    // Redis database number
}

// LoadConfig loads configuration from environment variables or uses defaults.
func LoadConfig() *Config {
	return &Config{
		MySQLDSN:      getEnv("MYSQL_DSN", "root:root@tcp(127.0.0.1:3306)/choicetech"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       0,
	}
}

// getEnv returns the value of the environment variable if set, otherwise returns the fallback value.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
