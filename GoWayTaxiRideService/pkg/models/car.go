package models

type Car struct {
	Id      int    `json:"id"`
	CarMake string `json:"car_make"`
	Color   string `json:"color"`
}
