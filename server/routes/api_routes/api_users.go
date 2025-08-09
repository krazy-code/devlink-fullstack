package api_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/controllers"
)

func UsersRoutes(r fiber.Router) {
	r.Get("/users", controllers.GetUsers)
}
