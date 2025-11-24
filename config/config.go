package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	JWTSecret  string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", ""),
		DBName:    getEnv("DB_NAME", ""),
		Port:      getEnv("PORT", "9999"),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
