package service

import (
	"RideService/internal/kafka"
	"RideService/pkg/database"
	"RideService/pkg/models"
	"RideService/pkg/models/request"
	"time"
)

func startOrder(order request.RequestOrder) {
	currentTime := time.Now()

	new_order := models.Order{
		UserId:         order.UserId,
		DriverId:       order.DriverId,
		TimeStartOrder: currentTime.Format("2006-01-02 15:04:05"),
		CarId:          order.CarId,
		OrderClass:     order.OrderClass,
		City:           order.City,
		StartPosition:  order.StartPosition,
		EndPosition:    order.EndPosition,
		Status:         "open",
	}

	database.DB.Model(&models.Driver{}).Where("id = ?", order.DriverId).Update("status", "busy")
	database.DB.Create(&new_order)
}

func endOrder(orderId int) {
	kafka.InitKafka()

	var reqPrice request.RequestPrice
	var reqOrder request.RequestOrder

	currentTime := time.Now()

	database.DB.Model(&models.Order{}).
		Where("id = ?", orderId).
		Updates(models.Order{
			TimeEndOrder: currentTime.Format("2006-01-02 15:04:05"),
			Status:       "completed",
			Distance:     reqOrder.Distance,
		})

	kafka.SendMessagePricing(reqOrder, "pricing-topic")

	database.DB.Model(&models.Order{}).
		Where("id = ?", orderId).
		Updates(models.Order{
			Price: reqPrice.Price,
		})

	database.DB.Model(&models.Driver{}).Where("id = ?", orderId).Update("status", "available")
}
