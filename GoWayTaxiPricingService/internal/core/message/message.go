package message

import (
	"GoWayTaxiPricingService/internal/core/service"
	"GoWayTaxiPricingService/pkg/models"
	"encoding/json"
	"fmt"
)

func ProcessMessage(data []byte) (int, float64, error) {
	fmt.Println("Consumer started")

	var orderMath models.OrderMath
	if err := json.Unmarshal(data, &orderMath); err != nil {
		return 0, 0, err
	}

	result := service.PricingOrder(orderMath)
	return orderMath.Id, result, nil
}
