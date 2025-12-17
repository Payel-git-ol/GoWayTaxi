package main

import (
	"RideService/internal/kafka"
	"RideService/internal/service"
	"RideService/internal/service/saving"
	"RideService/pkg/database"
	"RideService/pkg/models/request"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"strconv"
	"sync"
)

func main() {
	database.InitDB()
	saving.SaveCar()

	var wg sync.WaitGroup
	wg.Add(1)
	go kafka.GetMessageUserAndDriver(&wg)
	go kafka.GetMessageResultPricing(&wg)

	app := fiber.New()

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

	app.Post("/stop/taxi/order/:orderId", func(c fiber.Ctx) error {
		orderIdStr := c.Params("orderId")

		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("invalid order Id")
		}

		service.CreateOrderEnd(orderId)

		return c.JSON(fiber.Map{
			"order": "stopped",
		})
	})

	app.Post("/grade/order/:orderId", func(c fiber.Ctx) error {
		orderIdStr := c.Params("orderId")
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("invalid order Id")
		}

		body := c.Body()
		var grade request.GradeRequest
		errJson := json.Unmarshal(body, &grade)

		if errJson != nil {
			return c.Status(fiber.StatusBadRequest).SendString(errJson.Error())
		}

		service.GiveRating(orderId, grade)

		return c.JSON(fiber.Map{
			"Status": fiber.StatusOK,
		})
	})

	app.Listen(":5000")
}
