package utils

import "github.com/gofiber/fiber/v2"

func BadRequest(c *fiber.Ctx,message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": message,
	})
}

func Unuthorize(c *fiber.Ctx,message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": message,
	})
}

func Forbidden(c *fiber.Ctx,message string) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error": message,
	})
}

func Conflict(c *fiber.Ctx,message string) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"error": message,
	})
}

func NotFound(c *fiber.Ctx,message string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": message,
	})
}

func InternalServerError(c *fiber.Ctx,message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": message,
	})
}