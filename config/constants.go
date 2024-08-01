package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SERVER_URL string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	SERVER_URL = os.Getenv("SERVER_URL")
}
