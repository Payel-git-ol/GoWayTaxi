package message

import (
	"GoWayTaxiPricingService/internal/service"
	"GoWayTaxiPricingService/pkg/models"
	"encoding/json"
	"fmt"
)

func ProcessMessage(data []byte) (string, error) {
	fmt.Printf("Consumer started")

	var orderMath models.OrderMath

	if err := json.Unmarshal(data, &orderMath); err != nil {
		return "", err
	}

	service.PricingOrder(orderMath)

	return "Consumer finished", nil
}
