package service

import (
	"GoWayTaxiPricingService/pkg/models"
	"math"
	"strings"
	"time"
)

func PricingOrder(orderMath models.OrderMath) float64 {
	layout := "2006-01-02 15:04:05"
	start, _ := time.Parse(layout, orderMath.TimeStartOrder)
	end, _ := time.Parse(layout, orderMath.TimeEndOrder)
	durationMinutes := end.Sub(start).Minutes()

	distanceKm := orderMath.Distance

	multiplier := 1.0
	switch strings.ToLower(orderMath.OrderClass) {
	case "economy":
		multiplier = 1.5
	case "comfort":
		multiplier = 2.0
	case "business":
		multiplier = 2.5
	}

	baseFare := 3.0
	pricePerKm := 0.5
	pricePerMinute := 0.2

	total := (baseFare + distanceKm*pricePerKm + durationMinutes*pricePerMinute) * multiplier
	return math.Round(total*100) / 100
}
