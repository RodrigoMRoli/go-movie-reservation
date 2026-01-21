package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	log.Println("DB Host:", dbHost)
	log.Println("DB Port:", dbPort)
	log.Println("DB User:", dbUser)
	log.Println("DB Password:", dbPassword)
	log.Println("DB Name:", dbName)
}
