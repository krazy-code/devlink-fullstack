package routes

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/krazy-code/devlink/routes/api/v1/auth"
	users "github.com/krazy-code/devlink/routes/api/v1/users"
)

var (
	Prefix string
)

func InitRouter(a *fiber.App) {
	Prefix := "/api/v1"
	users.UsersRoutes(a, Prefix)
	auth.AuthRoutes(a, Prefix)

	a.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
	})
}
