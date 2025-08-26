package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Environment string
	Port        string
	AppName     string
	
	Database DatabaseConfig
	JWT      JWTConfig
	
	LogLevel  string
	LogFormat string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	URL      string 
}

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

func Load() *Config {
	return &Config{
		Environment: getEnv("ENV", "development"),
		Port:        getEnv("PORT", "8080"),
		AppName:     getEnv("APP_NAME", "turning-back"),
		
		Database: DatabaseConfig{
			URL:      getEnv("DATABASE_URL", ""),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "turning_back_user"),
			Password: getEnv("DB_PASSWORD", "turning_back_pass"),
			Name:     getEnv("DB_NAME", "turning_back_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
			Expiry: getEnvAsDuration("JWT_EXPIRY", "24h"),
		},
		
		LogLevel:  getEnv("LOG_LEVEL", "info"),
		LogFormat: getEnv("LOG_FORMAT", "json"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue string) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	if duration, err := time.ParseDuration(defaultValue); err == nil {
		return duration
	}
	return 24 * time.Hour
}