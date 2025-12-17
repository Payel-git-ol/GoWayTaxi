package main

import (
	"RideService/internal/kafka"
	"RideService/internal/service"
	"RideService/pkg/database"
	"RideService/pkg/models/request"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"strconv"
	"sync"
)

func main() {
	database.InitDB()
	service.SaveCar()

	var wg sync.WaitGroup
	wg.Add(1)
	go kafka.GetMessageUserAndDriver(&wg)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app": "GoWayTaxiRide",
		})
	})

	app.Post("/create/taxi/order", func(c fiber.Ctx) error {
		body := c.Body()
		var order request.RequestOrder
		err := json.Unmarshal(body, &order)

		if err != nil {
			return err
		}

		service.CreateOrderStart(order)

		return c.JSON(fiber.Map{
			"order": "run",
		})
	})

	app.Post("/stop/taxi/order", func(c fiber.Ctx) error {
		orderIdStr := c.Params("orderId")
		orderPriceStr := c.Params("orderPrice")

		orderPrice, err := strconv.ParseFloat(orderPriceStr, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("invalid order price")
		}

		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("invalid order Id")
		}

		service.CreateOrderEnd(orderId, orderPrice)

		return c.JSON(fiber.Map{
			"order": "stopped",
		})
	})

	app.Listen(":5000")
}
