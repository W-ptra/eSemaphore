package service

import (
	"eSemaphore-backend/database"
	"eSemaphore-backend/model"
	"github.com/gofiber/fiber/v2"
)

type Service struct {
	db 				*database.Database
	jwtKey 			[]byte
	jwtExpired		int

	chatWebSocket	ChatWebSocket
}

func CreateService(db *database.Database, jwtSecret []byte, jwtExpired int) Service {
	chatWebSocket := getChatWebSocket()
	return Service{
		db:db,
		jwtKey: jwtSecret,
		jwtExpired: jwtExpired,
		chatWebSocket: *chatWebSocket,
	}
}

func getUserFromContext(c *fiber.Ctx) model.User {
	user,_ := c.Locals("user").(model.User)
	return user
}