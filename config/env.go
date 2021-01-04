package config

import (
	"fmt"
	"os"
)

var (
	PORT    = getEnv("PORT", "8080")
	DB_HOST = getEnv("DB_HOST", "localhost")
	DB_USER = getEnv("DB_USER", "postgres")
	DB_PASS = getEnv("DB_PASS", "admin")
	DB_NAME = getEnv("DB_NAME", "spacehax")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
