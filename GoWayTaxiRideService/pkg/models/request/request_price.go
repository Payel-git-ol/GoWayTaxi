package request

type RequestPrice struct {
	OrderId int     `json:"order_id"`
	Price   float64 `json:"price"`
}
