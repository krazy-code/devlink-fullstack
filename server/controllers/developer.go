package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/models"
	"github.com/krazy-code/devlink/utils"
)

type developer struct {
	queries *database.Queries
}

func NewDeveloper(db *database.Queries) developer {
	return developer{
		queries: db,
	}
}

func (controllers *developer) Route(r fiber.Router) {
	const prefix = "/developers"
	r.Get(prefix, controllers.GetDevelopers)
	r.Get(prefix+"/:id", controllers.GetDeveloper)
	r.Post(prefix, controllers.CreateDeveloper)
	r.Put(prefix+"/:id", controllers.UpdateDeveloper)
	r.Delete(prefix+"/:id", controllers.DeleteDeveloper)
}

func (controllers *developer) GetDevelopers(c *fiber.Ctx) error {
	developers, err := controllers.queries.GetDevelopers()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data: fiber.Map{
				"count":      0,
				"developers": nil,
			},
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"count":      len(developers),
			"developers": developers,
		},
	})
}

func (controllers *developer) GetDeveloper(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	developer, err := controllers.queries.GetDeveloper(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: developer,
	})
}

func (controllers *developer) CreateDeveloper(c *fiber.Ctx) error {
	var req models.Developer

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}
	developerId, err := controllers.queries.CreateDeveloper(&req)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data: fiber.Map{
				"developer_id": developerId,
			},
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusCreated,
		Data: fiber.Map{
			"developer_id": developerId,
		},
	})
}

func (controllers *developer) UpdateDeveloper(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	foundedDeveloper, err := controllers.queries.GetDeveloper(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	req := &models.Developer{}
	if err := c.BodyParser(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}

	if err := controllers.queries.UpdateDeveloper(foundedDeveloper.Id, req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}

func (controllers *developer) DeleteDeveloper(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	foundedDeveloper, err := controllers.queries.GetDeveloper(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	req := &models.Developer{}
	if err := c.BodyParser(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}

	if err := controllers.queries.DeleteDeveloper(foundedDeveloper.Id); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}
