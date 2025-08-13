package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/middleware"
	api_routes "github.com/krazy-code/devlink/routes/api_routes"
)

var (
	Prefix string
)

func InitRouter(a *fiber.App) {
	middleware.Middlewares(a)

	api := a.Group("/api")

	v1 := api.Group("/v1", middleware.JWTProtected)

	api_routes.UsersRoutes(v1)
	api_routes.AuthRoutes(v1)
	api_routes.DevelopersRoutes(v1)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
	})
}
