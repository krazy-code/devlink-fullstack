package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/middleware"
	v1 "github.com/krazy-code/devlink/routes/api/v1"
)

var (
	Prefix string
)

func InitRouter(a *fiber.App) {
	// Middlewares
	middleware.Middlewares(a)
	// Group API v1
	Prefix := "/api/v1"
	v1.UsersRoutes(a, Prefix)
	v1.AuthRoutes(a, Prefix)
	v1.DevelopersRoutes(a, Prefix)

	a.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
	})
}
