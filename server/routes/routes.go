package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/controllers"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/middleware"
)

var (
	Prefix string
)

func InitRouter(a *fiber.App) {
	middleware.Middlewares(a)

	db, err := database.OpenDBConnection()
	if err != nil {
		panic(err)
	}

	authController := controllers.NewAuth(db)
	developerController := controllers.NewDeveloper(db)
	userController := controllers.NewUser(db)
	projectController := controllers.NewProject(db)

	api := a.Group("/api")

	v1 := api.Group("/v1", middleware.JWTProtected)

	userController.Route(v1)
	authController.Route(v1)
	developerController.Route(v1)
	projectController.Route(v1)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
	})
}
