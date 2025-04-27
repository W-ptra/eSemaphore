package router

import (
	"eSemaphore-backend/service"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(
	path string,
	api fiber.Router, 
	service service.Service,
){
	auth := api.Group(path)
	auth.Post("/login", service.Login)
	auth.Post("/register", service.Register)
}