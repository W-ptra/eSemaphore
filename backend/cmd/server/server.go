package main

import (
	"eSemaphore-backend/router"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")
	router.AuthRouter(api)
	router.ChatRouter(api)

	port := os.Getenv("PORT")
	listen := fmt.Sprintf("0.0.0.0:%v",port)

	app.Listen(listen)
}