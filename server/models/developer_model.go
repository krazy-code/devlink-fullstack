package models

type Developer struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Website  string `json:"website"`
	Github   string `json:"github"`
}
