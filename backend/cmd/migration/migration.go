package main

import (
	"eSemaphore-backend/config"
	"eSemaphore-backend/database"
	"eSemaphore-backend/model"
	"log"
)

func main() {
	config := config.GetConfig()

	db, err := database.NewDatabase(
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
	)
	if err!=nil{
		log.Fatal("Database connection failed",err)
	}
	log.Println("Establish database connection success")

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