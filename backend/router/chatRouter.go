package router

import (
	"eSemaphore-backend/middleware"
	"eSemaphore-backend/service"

	"github.com/gofiber/fiber/v2"
)

func ChatRouter(
	api fiber.Router,
	middleware middleware.Middleware,
	service service.Service,
){
	chat := api.Group("/chat")
	chat.Get("/",middleware.Authentication,service.GetChat)
}