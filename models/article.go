package models

type Article struct {
	BaseModel
	Title string `json:"title"`
	Body  string `json:"body"`
}
