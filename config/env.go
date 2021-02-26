package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	PORT    = getEnv("PORT")
	DB_HOST = getEnv("DB_HOST")
	DB_USER = getEnv("DB_USER")
	DB_PASS = getEnv("DB_PASS")
	DB_NAME = getEnv("DB_NAME")
	SECRET  = getEnv("SECRET")
)

func getEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
