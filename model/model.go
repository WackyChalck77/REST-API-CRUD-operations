package model

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
