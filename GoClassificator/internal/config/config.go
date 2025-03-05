package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/****************************************************
 * 													*
 * Configuration management functions				*
 * 													*
 ****************************************************/

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file, using system environment variables")
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}
