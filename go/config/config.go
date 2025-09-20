package config

import "os"

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	SSLMode    string
}

func LoadConfig() Config {
	return Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "1532"),
		DBName:     getEnv("DB_NAME", "postgres"),
		DBPort:     getEnv("DB_PORT", "5432"),
		SSLMode:    getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
