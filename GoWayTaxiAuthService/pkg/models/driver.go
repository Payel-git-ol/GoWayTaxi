package models

type Driver struct {
	Id       int    `json:"id"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
