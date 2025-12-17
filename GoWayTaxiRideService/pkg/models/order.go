package models

type Order struct {
	Id             int    `json:"id"`
	UserId         int    `json:"user_id"`
	DriverId       int    `json:"driver_id"`
	CarId          int    `json:"car_id"`
	TimeNow        string `json:"time"`
	TimeStartOrder string `json:"time_start_order"`
	TimeEndOrder   string `json:"time_end_order"`
	Route          string `json:"route"`
	OrderClass     string `json:"class"`
	Price          string `json:"price"`
}
