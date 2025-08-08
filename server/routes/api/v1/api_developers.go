package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/controllers"
)

func DevelopersRoutes(a *fiber.App, p string) {
	const prefix = "/developers"
	apiV1 := a.Group(p)
	apiV1.Get(prefix, controllers.GetDevelopers)
	apiV1.Get(prefix+"/:id", controllers.GetDeveloper)
	apiV1.Post(prefix, controllers.CreateDeveloper)
	apiV1.Put(prefix+"/:id", controllers.UpdateDeveloper)
}
