package models

type Skill struct {
	Id          int    `json:"id"`
	DeveloperId int    `json:"developer_id"`
	Name        string `json:"name"`
}

// - id (uuid)
// - developer_id (foreign key -> Developer.id)
// - name (string)
