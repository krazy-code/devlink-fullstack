package api_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/controllers"
)

func UsersRoutes(r fiber.Router) {
	const prefix = "/users"
	r.Get(prefix, controllers.GetUsers)
	r.Get(prefix+"/:id", controllers.GetUser)
	r.Post(prefix, controllers.CreateUser)
	r.Put(prefix+"/:id", controllers.UpdateUser)
	r.Delete(prefix+"/:id", controllers.DeleteUser)
}
