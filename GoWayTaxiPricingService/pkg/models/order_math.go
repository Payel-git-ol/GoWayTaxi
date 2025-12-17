package models

type OrderMath struct {
	Id             int     `json:"id"`
	TimeStartOrder string  `json:"time_start_order"`
	TimeEndOrder   string  `json:"time_end_order"`
	StartPosition  string  `json:"start_position"`
	EndPosition    string  `json:"end_position"`
	OrderClass     string  `json:"class"`
	Distance       float64 `json:"distance"`
}
