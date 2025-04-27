package service

import (
	//"eSemaphore-backend/dto"
	"eSemaphore-backend/utils"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	//"github.com/gofiber/websocket/v2"
)

func (s *Service) GetChats(c *fiber.Ctx) error {
	user := getUserFromContext(c)

	chats,err := s.db.GetChatByUserId(user.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound){
		return utils.NotFound(c,
			fmt.Sprintf(
				"No chats found with user id %s",
				user.Id,
			),
		)
	} else if err != nil{
		return utils.InternalServerError(c,"can't fetch chats")
	}

	// chatDtos := dto.ConvertChatModelsToDtos(
	// 	chats,
	// 	"test",
	// 	"test",
	// 	"ppp",
	// )

	return c.Status(200).JSON(fiber.Map{
		"data":chats,
	})
}

func (s *Service) GetChatLines(c *fiber.Ctx) error {
	//chatId := c.Params("chatId")


	user := getUserFromContext(c)

	chats,err := s.db.GetChatLinesByChatId(user.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound){
		return utils.NotFound(c,
			fmt.Sprintf(
				"No chats found with user id %s",
				user.Id,
			),
		)
	} else if err != nil{
		return utils.InternalServerError(c,"can't fetch chats")
	}

	// chatDtos := dto.ConvertChatModelsToDtos(
	// 	chats,
	// 	"test",
	// 	"test",
	// 	"ppp",
	// )

	return c.Status(200).JSON(fiber.Map{
		"data":chats,
	})
}