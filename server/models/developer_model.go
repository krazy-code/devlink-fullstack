package models

import "github.com/google/uuid"

type Developer struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Location  string    `json:"location"`
	Website   string    `json:"website"`
	Github    string    `json:"github"`
	CreatedAt string    `json:"created_at"`
}
