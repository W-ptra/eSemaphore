package router

import (
	"eSemaphore-backend/service"
	"github.com/gofiber/fiber/v2"
)

func ChatRouter(api fiber.Router){
	chat := api.Group("/chat")
	chat.Get("/",service.GetChat)
}