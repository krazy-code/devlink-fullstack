package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/krazy-code/devlink/routes/api/v1/users"
)

var (
	Prefix string
)

func InitRouter(a *fiber.App) {
	Prefix := "/api/v1"
	v1.UsersRoutes(a, Prefix)
}
