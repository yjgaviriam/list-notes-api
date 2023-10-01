package models

type Note struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostNote struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
