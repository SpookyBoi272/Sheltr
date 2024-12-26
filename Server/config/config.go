package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	Port string
	DBPort string
	DBHost string
	DBUser string
	DBPassword string
	DBName string
	JwtKey string
}

func Load() *Config{
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	return &Config{
		Port: getEnv("PORT", "8080"),
		DBHost: getEnv("DB_HOST","127.0.0.1"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName: getEnv("DB_NAME", "rent"),
		JwtKey: getEnv("JWT_KEY", ""),
	}
}

func getEnv(key string, fallback string) string{
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}