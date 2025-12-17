package service

import (
	"RideService/pkg/database"
	"RideService/pkg/models"
	"RideService/pkg/models/request"
)

func GiveRating(orderId int, req request.GradeRequest) {
	grade := models.GradeOrder{
		OrderId:  orderId,
		UserId:   req.UserId,
		DriverId: req.DriverId,
		Class:    req.Class,
		Grade:    req.Grade,
		Title:    req.Title,
	}

	database.DB.Create(&grade)
}
