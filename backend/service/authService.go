package service

import(
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message":"Login",
	})
}

func Register(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message":"Register",
	})
}