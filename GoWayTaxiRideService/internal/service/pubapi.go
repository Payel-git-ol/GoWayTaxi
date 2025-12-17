package service

import (
	"RideService/pkg/models/request"
)

func CreateOrderStart(order request.RequestOrder) {
	startOrder(order)
}

func CreateOrderEnd(orderId int) {
	endOrder(orderId)
}
