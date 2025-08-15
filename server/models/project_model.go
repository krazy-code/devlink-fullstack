package models

import "github.com/google/uuid"

type Project struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Techstacks  string    `json:"tech_stacks"`
	ProjectURL  string    `json:"project_url"`
	CreatedAt   string    `json:"created_at"`
}
