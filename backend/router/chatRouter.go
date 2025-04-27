package router

import (
	"eSemaphore-backend/middleware"
	"eSemaphore-backend/service"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/websocket/v2"
)

func ChatRouter(
	apiPath string,
	wsPath string,
	api fiber.Router,
	ws fiber.Router,
	middleware middleware.Middleware,
	service service.Service,
){
	chatApi := api.Group(apiPath)
	//chatWs := ws.Group(wsPath)

	chatApi.Get("/",)
	//chatWs.Get("/", websocket.New(service.ChatWebsocket))
}