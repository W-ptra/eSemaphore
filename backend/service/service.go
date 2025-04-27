package service

import (
	"eSemaphore-backend/database"
	"eSemaphore-backend/model"
	"github.com/gofiber/fiber/v2"
)

type Service struct {
	db 			*database.Database
	jwtKey 		[]byte
	jwtExpired	int
}

func CreateService(db *database.Database, jwtSecret []byte, jwtExpired int) Service {
	return Service{
		db:db,
		jwtKey: jwtSecret,
		jwtExpired: jwtExpired,
	}
}

func getUserFromContext(c *fiber.Ctx) model.User {
	user,_ := c.Locals("user").(model.User)
	return user
}