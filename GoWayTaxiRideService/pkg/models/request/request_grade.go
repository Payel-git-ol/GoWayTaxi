package request

type GradeRequest struct {
	OrderId  int    `json:"order_id"`
	UserId   int    `json:"user_id"`
	DriverId int    `json:"driver_id"`
	Class    string `json:"class"`
	Grade    int    `json:"grade"`
	Title    string `json:"title"`
}
