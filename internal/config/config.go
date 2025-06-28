package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func GetOuraToken() string {
	token := os.Getenv("OURA_TOKEN")
	if token == "" {
		log.Fatal("OURA_TOKEN environment variable is not set")
	}
	return token
}
