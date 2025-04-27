package service

import (
	"eSemaphore-backend/dto"
	"errors"
	"fmt"

	"eSemaphore-backend/model"
	"eSemaphore-backend/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (s *Service) Login(c *fiber.Ctx) error {

	var body dto.LoginRequest
	if err := c.BodyParser(&body); err != nil{
		return utils.BadRequest(c,"Invalid json")
	}

	isPropertyEmpty := (len(body.Email) == 0) || (len(body.Password) == 0)
	if isPropertyEmpty{
		return utils.BadRequest(c,"email or password can't empty")
	}

	isPasswordInvalid := len(body.Password) < 8 || len(body.Password) > 16
	if isPasswordInvalid{
		return utils.BadRequest(c,"password length must between 8 - 16")
	}

	user,err := s.db.GetUserByEmail(body.Email)
	if err != nil && errors.Is(err,gorm.ErrRecordNotFound){
		return utils.NotFound(c,
			fmt.Sprintf("User with email %v is not found",body.Email),
		)
	} else if err != nil{
		return utils.InternalServerError(c,"Something went wrong, when fetch user")
	}

	isPasswordNotMatch := !utils.ComparePassword(body.Password,user.Password)
	if isPasswordNotMatch{
		return utils.Unuthorize(c,"wrong password")
	}

	token := utils.CreateJwt(user,s.jwtKey,s.jwtExpired)

	return c.Status(200).JSON(fiber.Map{
		"token":token,
	})
}

func (s *Service) Register(c *fiber.Ctx) error {
	var body dto.RegisterRequest
	if err := c.BodyParser(&body); err != nil{
		return utils.BadRequest(c,"Invalid json")
	}

	isPropertyEmpty := (len(body.Name) == 0) || (len(body.Email) == 0) || (len(body.Password) == 0)
	if isPropertyEmpty{
		return utils.BadRequest(c,"name, email or password can't empty")
	}

	isPasswordInvalid := len(body.Password) < 8 || len(body.Password) > 16
	if isPasswordInvalid{
		return utils.BadRequest(c,"password length must between 8 - 16")
	}

	_,err :=  s.db.GetUserByEmail(body.Email)
	if err == nil{
		return utils.Conflict(c,"email already used")
	}  else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.InternalServerError(c, "Something went wrong, when do email checking")
	}

	newUser := model.User{
		Id: utils.GenerateUuid(),
		Name: body.Name,
		Email: body.Email,
		Password: utils.HashPassword(body.Password),
	}

	newUser,err = s.db.CreateUser(newUser)
	if err != nil{
		return utils.InternalServerError(c, "Something went wrong, when creating new user")
	}

	return c.Status(200).JSON(fiber.Map{
		"message":fmt.Sprintf(
			"sucessfully create new user with id %s",
			newUser.Id,
		),
	})
}