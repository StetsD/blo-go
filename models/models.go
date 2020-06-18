package models

type Article struct {
	ID int `json:"id"`
	Title string `json:"title'"`
	Content string `json:"content"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}