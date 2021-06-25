package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err loading .env file")
	}
	secretKey := os.Getenv(key)
	return secretKey
}