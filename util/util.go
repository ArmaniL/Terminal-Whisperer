package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(label string) string {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(label)

}
