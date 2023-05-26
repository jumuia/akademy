package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jumuia/akademy/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbURL := os.Getenv("DATABASE_URL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	cfg := server.Config{
		Port:       port,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBURL:      dbURL,
		DBHost:     dbHost,
		DBPort:     dbPort,
	}
	err = server.Run(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exiting the server...")
}
