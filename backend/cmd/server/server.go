package main

import (
	"eSemaphore-backend/database"
	"eSemaphore-backend/middleware"
	"eSemaphore-backend/router"
	"eSemaphore-backend/service"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main(){
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
	if err!=nil{
		log.Fatal("Database connection failed",err)
	}
	log.Println("Establish database connection success")

	jwtExpired,err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		log.Fatal("Can't load parse JWT_EXPIRED to int")
	}
	log.Println("Load jwtExpired success")

	service := service.CreateService(
		db,
		os.Getenv("JWT_SECRET"),
		jwtExpired,
	)
	middleware := middleware.GetMiddleware(
		os.Getenv("JWT_SECRET"),
	)

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")
	router.AuthRouter(api,service)
	router.ChatRouter(api,middleware,service)

	port := os.Getenv("PORT")
	listen := fmt.Sprintf("0.0.0.0:%v",port)

	app.Listen(listen)
}