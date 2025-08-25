package models

import "github.com/google/uuid"

type Follow struct {
	ID         uuid.UUID `json:"id"`
	FollowerID uuid.UUID `json:"follower_id"`
	FollowedID uuid.UUID `json:"followed_id"`
	CreatedAt  string    `json:"created_at"`
}
