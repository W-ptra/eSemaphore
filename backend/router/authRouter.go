package router

import (
	"eSemaphore-backend/service"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(api fiber.Router){
	auth := api.Group("/auth")
	auth.Post("/login",service.Login)
	auth.Post("/register",service.Register)
}