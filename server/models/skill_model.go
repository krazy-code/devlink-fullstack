package models

import "github.com/google/uuid"

type Skill struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
}
