package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/utils"
)

func GetUsers(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	users, err := db.GetUsers()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data: fiber.Map{
				"count": 0,
				"users": nil,
			},
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"count": len(users),
			"users": users,
		},
	})
}
