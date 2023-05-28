package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}

	// return the value based on a given key
	return os.Getenv(key)
}
