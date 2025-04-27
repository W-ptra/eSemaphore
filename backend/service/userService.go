package service

import (
	"eSemaphore-backend/dto"
	"eSemaphore-backend/utils"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var ALLOWED_EXTENSION =  map[string]bool{
    ".jpg":  true,
    ".jpeg": true,
    ".png":  true,
}

func (s *Service) GetUserProfile(c *fiber.Ctx) error {
	user := getUserFromContext(c)
	userProfileDto := dto.UserProfile{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
	}

	return c.Status(200).JSON(fiber.Map{
		"profile": userProfileDto,
	})
}

func (s *Service) UpdateUserProfile(c *fiber.Ctx) error {
	user := getUserFromContext(c)

	var body dto.UserProfile
	if err := c.BodyParser(&body); err != nil{
		return utils.BadRequest(c,"Invalid json")
	}

	isPropertyEmpty := len(body.Name) == 0
	if isPropertyEmpty{
		return utils.BadRequest(c,"name can't empty")
	}

	userById,_ := s.db.GetUserById(user.Id)

	userById.Name = body.Name
	userById.ImageUrl = body.ImageUrl

	if err := s.db.UpdateUserProfile(userById); err != nil{
		return utils.InternalServerError(c,"Something went wrong, can't update profile")
	}

	return c.Status(200).JSON(fiber.Map{
		"message":"successfully update profile",
	})
}

func (s *Service) UpdateUserEmail(c *fiber.Ctx) error {
	user := getUserFromContext(c)

	var body dto.UserProfile
	if err := c.BodyParser(&body); err != nil{
		return utils.BadRequest(c,"Invalid json")
	}

	isPropertyEmpty := len(body.Email) == 0
	if isPropertyEmpty{
		return utils.BadRequest(c,"email can't empty")
	}

	_,err := s.db.GetUserByEmail(body.Email)
	if err == nil{
		return utils.Conflict(c,"email already used")
	}  else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.InternalServerError(c, "Something went wrong, when do email checking")
	}

	userById,_ := s.db.GetUserById(user.Id)

	userById.Email = body.Email

	if err := s.db.UpdateUserEmail(userById); err != nil{
		return utils.InternalServerError(c,"Something went wrong, can't update email")
	}

	return c.Status(200).JSON(fiber.Map{
		"message":"successfully update email",
	})
}

func (s *Service) UpdateUserPassword(c *fiber.Ctx) error {
	user := getUserFromContext(c)

	var body dto.UserPassword
	if err := c.BodyParser(&body); err != nil{
		return utils.BadRequest(c,"Invalid json")
	}

	isPropertyEmpty := len(body.Password) == 0
	if isPropertyEmpty{
		return utils.BadRequest(c,"password can't empty")
	}

	isPasswordInvalid := len(body.Password) < 8 || len(body.Password) > 16
	if isPasswordInvalid{
		return utils.BadRequest(c,"password length must between 8 - 16")
	}

	userById,_ := s.db.GetUserById(user.Id)

	hashPassword := utils.HashPassword(body.Password)
	userById.Password = hashPassword

	if err := s.db.UpdateUserPassword(userById); err != nil{
		return utils.InternalServerError(c,"Something went wrong, can't update password")
	}

	return c.Status(200).JSON(fiber.Map{
		"message":"successfully update password",
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