package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (User) TableNameUser() string { return "users" }
