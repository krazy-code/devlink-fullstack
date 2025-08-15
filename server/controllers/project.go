package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/models"
	"github.com/krazy-code/devlink/utils"
)

type project struct {
	queries *database.Queries
}

func NewProject(db *database.Queries) project {
	return project{
		queries: db,
	}
}

func (controllers *project) Route(r fiber.Router) {
	const prefix = "/projects"
	r.Get(prefix, controllers.GetProjects)
	r.Get(prefix+"/:id", controllers.GetProject)
	r.Post(prefix, controllers.CreateProject)
	r.Put(prefix+"/:id", controllers.UpdateProject)
	r.Delete(prefix+"/:id", controllers.DeleteProject)
}

func (controllers *project) GetProjects(c *fiber.Ctx) error {
	projects, err := controllers.queries.GetProjects()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data: fiber.Map{
				"count":    0,
				"projects": nil,
			},
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"count":    len(projects),
			"projects": projects,
		},
	})
}

func (controllers *project) GetProject(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	project, err := controllers.queries.GetProject(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: project,
	})
}

func (controllers *project) CreateProject(c *fiber.Ctx) error {
	var req models.Project

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}
	projectId, err := controllers.queries.CreateProject(&req)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data: fiber.Map{
				"project_id": projectId,
			},
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusCreated,
		Data: fiber.Map{
			"project_id": projectId,
		},
	})
}

func (controllers *project) UpdateProject(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	foundedProject, err := controllers.queries.GetProject(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	req := &models.Project{}
	if err := c.BodyParser(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}

	if err := controllers.queries.UpdateProject(foundedProject.Id, req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}

func (controllers *project) DeleteProject(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	foundedProject, err := controllers.queries.GetProject(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	if err := controllers.queries.DeleteProject(foundedProject.Id); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusNoContent,
	})
}
