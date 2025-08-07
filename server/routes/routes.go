package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/middleware"
	auth "github.com/krazy-code/devlink/routes/api/v1/auth"
	users "github.com/krazy-code/devlink/routes/api/v1/users"
)

var (
	Prefix string
)

func InitRouter(a *fiber.App) {
	// Middlewares
	middleware.Middlewares(a)
	// Group API v1
	Prefix := "/api/v1"
	users.UsersRoutes(a, Prefix)
	auth.AuthRoutes(a, Prefix)

	a.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
	})
}
