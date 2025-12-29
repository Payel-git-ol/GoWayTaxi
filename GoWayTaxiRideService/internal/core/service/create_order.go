package service

import (
	kafka2 "RideService/internal/fetcher/kafka"
	"RideService/pkg/database"
	"RideService/pkg/models"
	"RideService/pkg/models/request"
	"fmt"
	"time"
)

func startOrder(order request.RequestOrder) {
	currentTime := time.Now()

	newOrder := models.Order{
		UserId:         order.UserId,
		DriverId:       order.DriverId,
		TimeStartOrder: currentTime.Format("2006-01-02 15:04:05"),
		CarId:          order.CarId,
		OrderClass:     order.OrderClass,
		City:           order.City,
		StartPosition:  order.StartPosition,
		EndPosition:    order.EndPosition,
		Status:         "open",
		Distance:       order.Distance,
	}

	database.DB.Model(&models.Driver{}).Where("id = ?", order.DriverId).Update("status", "busy")
	fmt.Println("Created order ID:", newOrder.Id)
	database.DB.Create(&newOrder)
}

func endOrder(orderId int) {
	kafka2.InitKafka()

	var order models.Order
	if err := database.DB.First(&order, orderId).Error; err != nil {
		return
	}

	currentTime := time.Now()

	database.DB.Model(&order).
		Updates(models.Order{
			TimeEndOrder: currentTime.Format("2006-01-02 15:04:05"),
			Status:       "completed",
		})

	reqOrder := request.RequestOrder{
		Id:             order.Id,
		UserId:         order.UserId,
		DriverId:       order.DriverId,
		CarId:          order.CarId,
		City:           order.City,
		OrderClass:     order.OrderClass,
		StartPosition:  order.StartPosition,
		EndPosition:    order.EndPosition,
		TimeStartOrder: order.TimeStartOrder,
		TimeEndOrder:   order.TimeEndOrder,
		Distance:       order.Distance,
	}

	kafka2.SendMessagePricing(reqOrder, "pricing-topic")

	var price request.RequestPrice

	database.DB.Model(&models.Driver{}).
		Where("id = ?", order.DriverId).
		Update("status", "available")

	database.DB.Model(&models.Order{}).
		Where("id = ?", order.Id).
		Update("price", price.Price)
}
