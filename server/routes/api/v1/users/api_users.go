package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/controllers"
)

func UsersRoutes(a *fiber.App, p string) {
	apiV1 := a.Group(p)
	apiV1.Get("/users", controllers.GetUsers)
}
