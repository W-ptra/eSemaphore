package main

import (
	"eSemaphore-backend/database"
	"eSemaphore-backend/model"
	"eSemaphore-backend/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loading env success")

	db, err := database.NewDatabase(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}
	log.Println("Establish database connection success")

	dummyUser1 := model.User{
		Id: utils.GenerateUuid(),
		Name: "jack Lim",
		Password: utils.HashPassword("12345678"),
		Email: "jackLimDummy@yahoo.com",
	}
	_,err = db.CreateUser(dummyUser1)
	if err != nil{
		log.Fatal("failed to created dummyUser1")
	}
	log.Println("create dummyUser1 success")

	dummyUser2 := model.User{
		Id: utils.GenerateUuid(),
		Name: "Joy Kristian",
		Password: utils.HashPassword("12345678"),
		Email: "kristianJ@onion.org",
	}
	_,err = db.CreateUser(dummyUser2)
	if err != nil{
		log.Fatal("failed to created dummyUser2")
	}
	log.Println("create dummyUser2 success")

	dummyChat := model.Chat{
		Id: utils.GenerateUuid(),
		Users: []model.User{dummyUser1,dummyUser2},
	}
	_,err = db.CreateChats(dummyChat)
	if err != nil{
		log.Fatal("failed to created dummyChat")
	}
	log.Println("create dummyChat success")

	dummyChatLine1 := model.ChatLine{
		Id: utils.GenerateUuid(),
		Content: "This is message from joy to jack",
		ChatID: dummyChat.Id,
		UserID: dummyUser2.Id,
	}
	_,err = db.CreateChatLine(dummyChatLine1)
	if err != nil{
		log.Fatal("failed to created dummyChatLine1")
	}
	log.Println("create CreateChatLine success")

	dummyChatLine2 := model.ChatLine{
		Id: utils.GenerateUuid(),
		Content: "This is message from jack to joy",
		ChatID: dummyChat.Id,
		UserID: dummyUser1.Id,
	}
	_,err = db.CreateChatLine(dummyChatLine2)
	if err != nil{
		log.Fatal("failed to created dummyChatLine2")
	}
	log.Println("create dummyChatLine2 success")

	log.Println("Seeder success")
}