package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not found", key)
	}
	return value, nil
}
