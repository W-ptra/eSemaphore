package service

import(
	"github.com/gofiber/fiber/v2"
)

func (s *Service) GetChat(c *fiber.Ctx) error {
	user := c.Locals("user")

	return c.Status(200).JSON(fiber.Map{
		"message":user,
	})
}
