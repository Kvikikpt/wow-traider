package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	port string
	databaseUrl string
}

func ReadConfig() *Config{
	//loading .env files
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		port: os.Getenv("PORT"),
		databaseUrl: os.Getenv("DATABASE_URL"),
	}
}
