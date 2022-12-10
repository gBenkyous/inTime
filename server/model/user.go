package model

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" example:"Jack"`
	Lastname  string `json:"lastname" example:"Packard"`
	Email     string `json:"email" example:"test@mail.com"`
	IsAdmin   bool   `json:"isAdmin"`
	AuthCount int    `json:"authCount"`
}
