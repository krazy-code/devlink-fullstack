package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/krazy-code/devlink/database"
	"github.com/krazy-code/devlink/utils"
)

type follow struct {
	queries *database.Queries
}

func NewFollowController(db *database.Queries) follow {
	return follow{queries: db}
}

func (controllers *follow) Route(r fiber.Router) {
	const prefix = "/follows"
	r.Post(prefix, controllers.CreateFollow)
	r.Delete(prefix, controllers.DeleteFollow)
	r.Get(prefix, controllers.ListFollowed)
}

// POST /follows
func (fc *follow) CreateFollow(c *fiber.Ctx) error {
	var req struct {
		FollowerID string `json:"follower_id"`
		FollowedID string `json:"followed_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusBadRequest, Errors: err.Error()})
	}
	log.Printf(`Follower: %s`, req.FollowerID)
	followerID, err := uuid.Parse(req.FollowerID)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusBadRequest, Errors: "Invalid follower_id"})
	}
	followedID, err := uuid.Parse(req.FollowedID)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusBadRequest, Errors: "Invalid followed_id"})
	}
	if err := fc.queries.CreateFollow(followerID, followedID); err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusInternalServerError, Errors: err.Error()})
	}
	return utils.ResponseParser(c, utils.Response{Code: fiber.StatusOK, Data: fiber.Map{"msg": "Followed successfully"}})
}

// DELETE /follows
func (fc *follow) DeleteFollow(c *fiber.Ctx) error {
	followerID := c.Query("follower_id")
	followedID := c.Query("followed_id")
	fID, err := uuid.Parse(followerID)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusBadRequest, Errors: "Invalid follower_id"})
	}
	fdID, err := uuid.Parse(followedID)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusBadRequest, Errors: "Invalid followed_id"})
	}
	if err := fc.queries.DeleteFollow(fID, fdID); err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusInternalServerError, Errors: err.Error()})
	}
	return utils.ResponseParser(c, utils.Response{Code: fiber.StatusOK, Data: fiber.Map{"msg": "Unfollowed successfully"}})
}

// GET /follows?follower_id=xxx
func (fc *follow) ListFollowed(c *fiber.Ctx) error {
	followerID := c.Query("follower_id")
	fID, err := uuid.Parse(followerID)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusBadRequest, Errors: "Invalid follower_id"})
	}
	follows, err := fc.queries.ListFollowed(fID)
	if err != nil {
		return utils.ResponseParser(c, utils.Response{Code: fiber.StatusInternalServerError, Errors: err.Error()})
	}
	return utils.ResponseParser(c, utils.Response{Code: fiber.StatusOK, Data: fiber.Map{"follows": follows}})
}
