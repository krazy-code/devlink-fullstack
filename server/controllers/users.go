package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/models"
	"github.com/krazy-code/devlink/utils"
)

type user struct {
	queries *database.Queries
}

func NewUser(db *database.Queries) user {
	return user{
		queries: db,
	}
}

func (controllers *user) Route(r fiber.Router) {
	const prefix = "/users"
	r.Get(prefix, controllers.GetUsers)
	r.Get(prefix+"/:id", controllers.GetUser)
	r.Post(prefix, controllers.CreateUser)
	r.Put(prefix+"/:id", controllers.UpdateUser)
	r.Delete(prefix+"/:id", controllers.DeleteUser)
}

func (controllers *user) GetUsers(c *fiber.Ctx) error {
	users, err := controllers.queries.GetUsers()
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

func (controllers *user) GetUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	user, err := controllers.queries.GetUser(id)
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

func (controllers *user) CreateUser(c *fiber.Ctx) error {
	var req models.User

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}
	userId, err := controllers.queries.CreateUser(&req)
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

func (controllers *user) UpdateUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	foundedUser, err := controllers.queries.GetUser(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	req := &models.User{}
	file, err := c.FormFile("avatar")
	if err == nil && file != nil {
		// Save the file or process as needed
		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save avatar"})
		}
		// Save file.Filename or path to DB as avatar
	}
	// if err := c.BodyParser(req); err != nil {
	// 	return utils.ResponseParser(c, utils.Response{
	// 		Code:   fiber.StatusBadRequest,
	// 		Errors: err.Error(),
	// 	})
	// }

	if err := controllers.queries.UpdateUser(foundedUser.Id, req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}

func (controllers *user) DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	foundedUser, err := controllers.queries.GetUser(id)
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

	if err := controllers.queries.DeleteUser(foundedUser.Id); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}
