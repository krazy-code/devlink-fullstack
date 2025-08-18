package middleware

import (
	"os"
	"slices"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/utils"
)

func JWTProtected(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		Filter: func(c *fiber.Ctx) bool {
			excludedPaths := []string{"/api/v1/auth", "/api/v1/auth/register", "/api/v1/auth/profile"}
			return slices.Contains(excludedPaths, c.Path())
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.ResponseParser(c, utils.Response{
				Code:   fiber.StatusUnauthorized,
				Errors: err.Error(),
			})
		},

		SuccessHandler: func(c *fiber.Ctx) error {
			authHeader := c.Get("Authorization")
			bearerToken := strings.Split(authHeader, " ")
			tokenString := bearerToken[1]

			if tokenString == "null" {
				return utils.ResponseParser(c, utils.Response{
					Code: fiber.StatusOK,
					Data: nil,
				})
			}

			claims, err := utils.VerifyToken(tokenString)
			if err != nil {
				return utils.ResponseParser(c, utils.Response{
					Code:   fiber.StatusUnauthorized,
					Errors: err.Error(),
				})
			}
			c.Locals("access_token_claims", claims)
			c.Next()
			return nil
		},
	})(c)
}
