package models

type Driver struct {
	Id         int    `json:"id"`
	DriverName string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Rating     string `json:"rating"`
	Status     string `json:"status"`
	CarId      int    `json:"car_id"`
}

func (Driver) TableNameDriver() string { return "drivers" }
