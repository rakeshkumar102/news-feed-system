package configs

import (
	"errors"
	"os"
)

func EnvMap(key string) (string, error) {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal("Error loading env :", err.Error())
	// }

	if os.Getenv(key) == "" {
		return "", errors.New("env not found")
	} else {
		return os.Getenv(key), nil
	}
}