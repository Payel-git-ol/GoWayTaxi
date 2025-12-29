package message

import (
	"RideService/pkg/database"
	"RideService/pkg/models"
	"RideService/pkg/models/request"
	"encoding/json"
)

func ProcessMessagePricing(data []byte) error {
	var msg request.RequestPrice
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}

	database.DB.Model(&models.Order{}).
		Where("id = ?", msg.OrderId).
		Update("price", msg.Price)

	return nil
}
