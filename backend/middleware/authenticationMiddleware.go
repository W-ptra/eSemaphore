package middleware

import (
	"eSemaphore-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authentication(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	if len(authorization) < 7{
		return utils.Unuthorize(c,"Missing authorization")
	}

	bearerToken := authorization[0:7]
	if bearerToken != "Bearer "{
		return utils.Unuthorize(c,"Missing Bearer token")
	}

	bearerToken = authorization[7:]
	user,err := utils.GetJwt(bearerToken,m.jwtSecret)
	
	if err != nil{
		return utils.Unuthorize(c,"token invalid or expired")
	}
	c.Locals("user",user)

	return c.Next()
}