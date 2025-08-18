package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt any       `json:"createdat"`
}

type UserDetail struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Bio         *string   `json:"bio"`
	GithubURL   *string   `json:"github_url"`
	WebsiteURL  *string   `json:"website_url"`
	LinkedinURL *string   `json:"linkedin_url"`
	Location    *string   `json:"location"`
	CreatedAt   any       `json:"createdat"`
}
