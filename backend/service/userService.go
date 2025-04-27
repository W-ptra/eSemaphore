package service

import (
	"eSemaphore-backend/dto"
	"eSemaphore-backend/model"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) getUserProfile(c *fiber.Ctx) error {
	user,_ := c.Locals("user").(model.User)
	userProfileDto := dto.UserProfile{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
	}

	return c.Status(200).JSON(fiber.Map{
		"profile": userProfileDto,
	})
}