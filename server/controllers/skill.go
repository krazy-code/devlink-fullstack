package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/models"
	"github.com/krazy-code/devlink/utils"
)

type skill struct {
	queries *database.Queries
}

func NewSkill(db *database.Queries) skill {
	return skill{
		queries: db,
	}
}

func (controllers *skill) Route(r fiber.Router) {
	const prefix = "/skills"
	r.Get(prefix, controllers.GetSkills)
	r.Get(prefix+"/:id", controllers.GetSkill)
	r.Post(prefix, controllers.CreateSkill)
	r.Put(prefix+"/:id", controllers.UpdateSkill)
	r.Delete(prefix+"/:id", controllers.DeleteSkill)
}

func (controllers *skill) GetSkills(c *fiber.Ctx) error {
	skills, err := controllers.queries.GetSkills()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data: fiber.Map{
				"count":  0,
				"skills": nil,
			},
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"count":  len(skills),
			"skills": skills,
		},
	})
}

func (controllers *skill) GetSkill(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	skill, err := controllers.queries.GetSkill(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: skill,
	})
}

func (controllers *skill) CreateSkill(c *fiber.Ctx) error {
	var req models.Skill

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}
	skillId, err := controllers.queries.CreateSkill(&req)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data: fiber.Map{
				"skill_id": skillId,
			},
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusCreated,
		Data: fiber.Map{
			"skill_id": skillId,
		},
	})
}

func (controllers *skill) UpdateSkill(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	foundedSkill, err := controllers.queries.GetSkill(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	req := &models.Skill{}
	if err := c.BodyParser(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}

	if err := controllers.queries.UpdateSkill(foundedSkill.Id, req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}

func (controllers *skill) DeleteSkill(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	foundedSkill, err := controllers.queries.GetSkill(id)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	if err := controllers.queries.DeleteSkill(foundedSkill.Id); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusNoContent,
	})
}
