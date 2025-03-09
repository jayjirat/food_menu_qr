package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIPort        string
	DB_HOST        string
	DB_USER        string
	DB_PORT        string
	DB_PASSWORD    string
	DB_NAME        string
	SECRET_KEY_JWT string
}

var AppConfig Config

func LoadConfig() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading config")
	}

	AppConfig = Config{
		APIPort:        GetEnv("PORT", "8080"),
		DB_HOST:        GetEnv("DB_HOST", "localhost"),
		DB_USER:        GetEnv("DB_USER", "postgres"),
		DB_PORT:        GetEnv("DB_PORT", "5432"),
		DB_PASSWORD:    GetEnv("DB_PASSWORD", "secret"),
		DB_NAME:        GetEnv("DB_NAME", "food-menu-qr"),
		SECRET_KEY_JWT: GetEnv("SECRET_KEY_JWT", "mysecret"),
	}
}

func GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
