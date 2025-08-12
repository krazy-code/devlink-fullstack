package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/models"
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

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	user, err := db.GetUser(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var req models.User

	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}
	userId, err := db.CreateUser(&req)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data: fiber.Map{
				"user_id": userId,
			},
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusCreated,
		Data: fiber.Map{
			"user_id": userId,
		},
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	foundedUser, err := db.GetUser(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	req := &models.User{}
	if err := c.BodyParser(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}

	if err := db.UpdateUser(foundedUser.Id, req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	foundedUser, err := db.GetUser(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	req := &models.User{}
	if err := c.BodyParser(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}

	if err := db.DeleteUser(foundedUser.Id); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}
