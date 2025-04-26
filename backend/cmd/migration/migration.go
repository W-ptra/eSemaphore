package main

import (
	"eSemaphore-backend/database"
	"eSemaphore-backend/model"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewDatabase(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	if err!=nil{
		log.Fatal("Database connection failed",err)
	}

	err = db.Migrate(
		&model.User{},
		&model.Chat{},
		&model.ChatLine{},
	)
	if err!=nil{
		log.Fatal("Migration failed",err)
	}

	log.Println("Migration success")
}