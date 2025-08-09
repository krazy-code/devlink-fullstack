package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/utils"
)

func AuthMidlleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	// token := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	if authHeader == "" {
		return utils.ResponseParser(c, utils.Response{
			Code: fiber.StatusUnauthorized,
			Errors: fiber.Map{
				"general": utils.GetStatusMessage(401),
			},
		})
	}
	return c.Next()
}
