package configs

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMap(key string) (string, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env")
	}

	if os.Getenv(key) == "" {
		return "", errors.New("env not found")
	} else {
		return os.Getenv(key), nil
	}
}