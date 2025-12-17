package models

type Order struct {
	Id             int     `json:"id"`
	UserId         int     `json:"user_id"`
	DriverId       int     `json:"driver_id"`
	CarId          int     `json:"car_id"`
	TimeStartOrder string  `json:"time_start_order"`
	TimeEndOrder   string  `json:"time_end_order"`
	City           string  `json:"city"`
	StartPosition  string  `json:"start_position"`
	EndPosition    string  `json:"end_position"`
	OrderClass     string  `json:"class"`
	Price          float64 `json:"price"`
	Status         string  `json:"status"`
	Distance       float64 `json:"distance"`
}
