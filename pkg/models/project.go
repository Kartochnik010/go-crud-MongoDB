package models

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerLogin  string `json:"owner_id"`
}
