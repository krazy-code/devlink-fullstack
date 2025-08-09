package api_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/controllers"
)

func DevelopersRoutes(r fiber.Router) {
	const prefix = "/developers"
	r.Get(prefix, controllers.GetDevelopers)
	r.Get(prefix+"/:id", controllers.GetDeveloper)
	r.Post(prefix, controllers.CreateDeveloper)
	r.Put(prefix+"/:id", controllers.UpdateDeveloper)
	r.Delete(prefix+"/:id", controllers.DeleteDeveloper)
}
