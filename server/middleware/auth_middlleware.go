package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/utils"
)

func JWTProtected(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		Filter: func(c *fiber.Ctx) bool {
			excludedPaths := []string{"/api/v1/auth", "/api/v1/auth/register"}
			for _, path := range excludedPaths {
				if c.Path() == path {
					return true
				}
			}
			return false
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.ResponseParser(c, utils.Response{
				Code:   fiber.StatusUnauthorized,
				Errors: err.Error(),
			})
		},
	})(c)
}
