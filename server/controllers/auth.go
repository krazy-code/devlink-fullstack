package controllers

import (
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/models"
	"github.com/krazy-code/devlink/utils"
)

type auth struct {
	queries *database.Queries
}

func NewAuth(db *database.Queries) auth {
	return auth{
		queries: db,
	}
}

func (controllers *auth) Route(r fiber.Router) {
	const prefix = "/auth"
	r.Post(prefix, controllers.PostLogin)
	r.Post(prefix+"/register", controllers.PostRegister)
	r.Post(prefix+"/logout", controllers.PostLogout)
	r.Get(prefix+"/profile", controllers.GetProfile)
}

func (controllers *auth) PostLogin(c *fiber.Ctx) error {
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

	userID, err := controllers.queries.PostLogin(&req)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}

	user, err := controllers.queries.GetUser(userID)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
			Data:   userID,
		})
	}
	claims := jwt.MapClaims{
		"name":    user.Name,
		"user_id": userID,
		"admin":   true,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := controllers.queries.UpdateAccessToken(userID, t); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})

	}
	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"user_id": userID,
			"token":   t,
			"user":    user,
		},
	})
}

func (controllers *auth) PostRegister(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	// var developerReq models.Developer
	// if err := c.BodyParser(&developerReq); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	// }

	if err := validator.New().Struct(req); err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusBadRequest,
			Errors: utils.ValidatorErrors(err),
		})
	}

	userID, err := controllers.queries.PostRegister(&req)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	// devId, err := controllers.queries.CreateDeveloper(&developerReq)
	// if err != nil {
	// 	return utils.ResponseParser(c, utils.Response{
	// 		Code:   fiber.StatusInternalServerError,
	// 		Errors: err.Error(),
	// 	})
	// }

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"user_id": userID,
			// "dev_id":  devId,
		},
	})
}

func (controllers *auth) PostLogout(c *fiber.Ctx) error {
	request := models.LogoutRequest{}
	request.AccessTokenClaims = c.Locals("access_token_claims").(jwt.MapClaims)
	userId, err := controllers.queries.PostLogout(request)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusInternalServerError,
			Errors: err.Error(),
		})
	}
	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"user_id": userId,
		},
	})
}

func (controllers *auth) GetProfile(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	tokenString := bearerToken[1]

	if tokenString == "null" {
		return utils.ResponseParser(c, utils.Response{
			Code: fiber.StatusOK,
			Data: nil,
		})
	}

	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusUnauthorized,
			Errors: err.Error(),
		})
	}

	userId, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}
	user, err := controllers.queries.GetUser(userId)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{
			Code:   fiber.StatusNotFound,
			Errors: err.Error(),
		})
	}

	return utils.ResponseParser(c, utils.Response{
		Code: fiber.StatusOK,
		Data: fiber.Map{
			"id":    user.Id,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
