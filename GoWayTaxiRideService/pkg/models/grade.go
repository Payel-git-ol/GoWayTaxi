package models

type GradeOrder struct {
	Id       int    `json:"id"`
	OrderId  int    `json:"order_id"`
	UserId   int    `json:"user_id"`
	DriverId int    `json:"driver_id"`
	Grade    int    `json:"grade"`
	Title    string `json:"title"`
}
