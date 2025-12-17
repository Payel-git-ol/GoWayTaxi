package service

import (
	"RideService/pkg/database"
	"RideService/pkg/models"
	"RideService/pkg/models/request"
	"time"
)

func CreateOrderStart(order request.RequestOrder) {
	currentTime := time.Now()

	new_order := models.Order{
		UserId:         order.UserId,
		DriverId:       order.DriverId,
		TimeStartOrder: currentTime.Format("2006-01-02 15:04:05"),
		Route:          order.Route,
		CarId:          order.CarId,
		OrderClass:     order.OrderClass,
	}

	database.DB.Create(&new_order)
}

func CreateOrderEnd(orderId int, price float64) {
	currentTime := time.Now()

	database.DB.Model(&models.Order{}).
		Where("id = ?", orderId).
		Updates(models.Order{
			TimeEndOrder: currentTime.Format("2006-01-02 15:04:05"),
			Price:        price,
			Status:       "completed",
		})
}
