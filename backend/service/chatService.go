package service

import(
	"github.com/gofiber/fiber/v2"
)

func GetChat(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message":"get chat",
	})
}
