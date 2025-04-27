package main

import (
	"eSemaphore-backend/config"
	"eSemaphore-backend/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main(){
	config := config.GetConfig()
	
	app := fiber.New()
	app.Use(logger.New())
	app.Static("/public","./public")

	app.Use(config.GetCorsConfig())

	router.CreateRouter(app,config)

	listen := fmt.Sprintf(
		"0.0.0.0:%v",
		config.PORT,
	)

	app.Listen(listen)
}