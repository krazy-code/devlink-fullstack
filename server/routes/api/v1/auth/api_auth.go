package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/controllers"
)

const prefix = "/auth"

func AuthRoutes(a *fiber.App, p string) {
	apiV1 := a.Group(p)
	apiV1.Post(prefix, controllers.PostLogin)
	apiV1.Post(prefix+"/register", controllers.PostRegister)
}
