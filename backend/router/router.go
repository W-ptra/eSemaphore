package router

import (
	"eSemaphore-backend/config"
	"eSemaphore-backend/database"
	"eSemaphore-backend/middleware"
	"eSemaphore-backend/service"
	"eSemaphore-backend/utils"
	"log"
	"github.com/gofiber/fiber/v2"
)

func CreateRouter(app *fiber.App, config config.Config) {
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

	service := service.CreateService(
		db,
		config.JWT_SECRET,
		config.JWT_EXPIRED,
	)
	log.Println("Create services success")

	middleware := middleware.GetMiddleware(
		config.JWT_SECRET,
	)
	log.Println("Create middlewares success")

	api := app.Group("/api")
	ws := app.Group("/ws") // ws = websocket

	AuthRouter("/auth", api, service)
	UserRouter("/user", api, middleware, service)
	ChatRouter("/chat","/ws", api, ws, middleware, service)

	app.Use(utils.NotFoundPage)
}