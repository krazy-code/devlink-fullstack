package api_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/controllers"
)

func AuthRoutes(r fiber.Router) {
	const prefix = "/auth"
	r.Post(prefix, controllers.PostLogin)
	r.Post(prefix+"/register", controllers.PostRegister)
	r.Post(prefix+"/logout", controllers.PostLogout)
}
