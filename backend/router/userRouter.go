package router

import (
	"eSemaphore-backend/middleware"
	"eSemaphore-backend/service"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(
	path string,
	api fiber.Router,
	middleware middleware.Middleware,
	service service.Service,
){
	user := api.Group(path)
	user.Get("/", middleware.Authentication, service.GetUserProfile)
	user.Post("/profile/image", middleware.Authentication, service.UploadUserProfilePicture)
}