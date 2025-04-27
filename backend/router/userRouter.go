package router

import (
	"eSemaphore-backend/middleware"
	"eSemaphore-backend/service"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(
	api fiber.Router,
	middleware middleware.Middleware,
	service service.Service,
){
	chat := api.Group("/user")
	chat.Get("/",middleware.Authentication,service.GetChat)
}