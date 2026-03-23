package models

type List struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
}
