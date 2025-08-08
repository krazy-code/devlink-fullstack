package models

type Developer struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Location  string `json:"location"`
	Website   string `json:"website"`
	Github    string `json:"github"`
	CreatedAt string `json:"created_at"`
}
