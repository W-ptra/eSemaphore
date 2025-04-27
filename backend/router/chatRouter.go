package router

import (
	"eSemaphore-backend/middleware"
	"eSemaphore-backend/service"

	"github.com/gofiber/fiber/v2"
)

func ChatRouter(
	path string,
	api fiber.Router,
	middleware middleware.Middleware,
	service service.Service,
){
	chat := api.Group(path)
	chat.Get("/", middleware.Authentication, service.GetChat)
}