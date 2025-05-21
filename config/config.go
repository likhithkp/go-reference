package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error while loading the .env file: %v", err)
		return ""
	}

	url := os.Getenv("POSTGRES_URL")
	if url == "" {
		log.Fatalln("No url found")
		return ""
	}

	return url
}
