package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/models"
	"github.com/krazy-code/devlink/utils"
)

func PostLogin(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}

	if err := validator.New().Struct(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: utils.ValidatorErrors(err),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	userID, err := db.PostLogin(&req)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"user_id": userID,
			"token":   "access_token",
		},
	})
}

func PostRegister(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	if err := validator.New().Struct(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: utils.ValidatorErrors(err),
		})
	}

	userID, err := db.PostRegister(&req)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"user_id": userID,
		},
	})
}

func PostLogout(c *fiber.Ctx) error {
	// var req models.RegisterRequest
	// db, err := database.OpenDBConnection()
	// if err != nil {
	// 	return utils.ResponseParser(c, utils.Response{
	// 		Code:   fiber.StatusInternalServerError,
	// 		Errors: err.Error(),
	// 	})
	// }

	// userID, err := db.PostRegister(&req)
	// if err != nil {
	// 	return utils.ResponseParser(c, utils.Response{
	// 		Code:   fiber.StatusInternalServerError,
	// 		Errors: err.Error(),
	// 	})
	// }
	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
	})
}
