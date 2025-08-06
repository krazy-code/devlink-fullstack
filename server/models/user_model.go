package models

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordhash"`
	CreatedAt    string `json:"createdat"`
}
