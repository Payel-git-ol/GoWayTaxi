package request

type RequestOrder struct {
	Id             int     `json:"id"`
	UserId         int     `json:"user_id"`
	DriverId       int     `json:"driver_id"`
	CarId          int     `json:"car_id"`
	TimeStartOrder string  `json:"time_start_order"`
	TimeEndOrder   string  `json:"time_end_order"`
	Route          string  `json:"route"`
	City           string  `json:"city"`
	OrderClass     string  `json:"class"`
	Price          float64 `json:"price"`
}
