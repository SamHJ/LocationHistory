package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Config func returns the specified environment variable based on supplied key
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
