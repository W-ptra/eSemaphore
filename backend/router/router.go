package router

import (
	"eSemaphore-backend/database"
	"eSemaphore-backend/middleware"
	"eSemaphore-backend/service"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func CreateRouter(app *fiber.App) {
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

	api := app.Group("/api")
	AuthRouter(api,service)
	ChatRouter(api,middleware,service)
}