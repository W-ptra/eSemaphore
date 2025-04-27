package service

import (
	"eSemaphore-backend/dto"
	"eSemaphore-backend/model"
	"eSemaphore-backend/utils"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

var ALLOWED_EXTENSION =  map[string]bool{
    ".jpg":  true,
    ".jpeg": true,
    ".png":  true,
}

func (s *Service) GetUserProfile(c *fiber.Ctx) error {
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

func (s *Service) UploadUserProfilePicture(c *fiber.Ctx) error {
	file,err := c.FormFile("image")
	if err != nil {
		log.Println(err)
		return utils.BadRequest(c,"Invalid form format")
	}

	fileSizeThreshold :=  10*1024*1024 // 10 MB
	isFileTooLarge := file.Size > int64(fileSizeThreshold)
	if isFileTooLarge {
		return utils.BadRequest(c,"Image size too large")
	}
	
	if isFileExtentionNotValid(file){
		return utils.BadRequest(c,"Invalid image format, allowed format is jpg, jpeg and png")
	}

	fileName := utils.GenerateUuid()
	savePath := fmt.Sprintf("./public/profile/%s",fileName)
	if err := c.SaveFile(file,savePath); err != nil {
		log.Println(err)
		return utils.InternalServerError(c,"Something went wrong, when saving profile image")
	}

	return c.Status(200).JSON(fiber.Map{
		"message":"successfully save profile image",
		"url": savePath[1:],
	})
}

func isFileExtentionNotValid(file *multipart.FileHeader) bool {
	extension := filepath.Ext(file.Filename)
	return !ALLOWED_EXTENSION[extension] // using hashmap for faster checking O(1)
}