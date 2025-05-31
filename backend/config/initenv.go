package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Не найден файл .env")
	}

	if DB_USERNAME := os.Getenv("DB_USERNAME"); DB_USERNAME == "" {
		os.Setenv("DB_USERNAME", "artisan")
	}

	if DB_PASSWORD := os.Getenv("DB_PASSWORD"); DB_PASSWORD == "" {
		os.Setenv("DB_PASSWORD", "artisan")
	}

	if DB_DATABASE := os.Getenv("DB_DATABASE"); DB_DATABASE == "" {
		os.Setenv("DB_DATABASE", "artisan")
	}

	if DB_PORT := os.Getenv("DB_PORT"); DB_PORT == "" {
		os.Setenv("DB_PORT", "5432")
	}

	if DB_ADRESS := os.Getenv("DB_ADDRESS"); DB_ADRESS == "" {
		os.Setenv("DB_ADRESS", "127.0.0.1")
	}
}
