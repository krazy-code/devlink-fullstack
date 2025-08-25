package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/models"
	"github.com/krazy-code/devlink/utils"
)

type testing struct {
	queries *database.Queries
}

func NewTestingController(db *database.Queries) testing {
	return testing{queries: db}

}
func (controllers *testing) Route(r fiber.Router) {
	r.Put("/testing", controllers.PutTesting)
}
func (controllers *testing) PutTesting(c *fiber.Ctx) error {
	var form models.User
	if err := c.BodyParser(&form); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: err.Error(),
		})
	}
	log.Printf("%v", form)

	file, err := c.FormFile("avatar")
	if err == nil && file != nil {
		// err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
		// if err != nil {
		// 	return utils.ResponseParser(c, utils.Response{
		// 		Code:   fiber.StatusInternalServerError,
		// 		Errors: err.Error(),
		// 	})
		// }
		// Save file.Filename or path to DB as avatar
	} else {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			// "name":      form.Name,
			// "bio":       form.Bio,
			// "email":     form.Email,
			"file_name": file.Filename,
		},
	})
}
