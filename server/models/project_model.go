package models

type Project struct {
	Id          int      `json:"id"`
	DeveloperId int      `json:"developer_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Techstacks  []string `json:"tech_stacks"`
	Link        string   `json:"link"`
}

// - id (uuid)
// - developer_id (foreign key -> Developer.id)
// - title (string)
// - description (text)
// - tech_stack (text or string[])
// - link (string)
