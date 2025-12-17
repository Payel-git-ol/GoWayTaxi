package main

import (
	"GoWayTaxiPricingService/internal/kafka"
	"github.com/gofiber/fiber/v3"
	"sync"
)

func main() {
	app := fiber.New()
	var wg sync.WaitGroup
	wg.Add(1)
	go kafka.GetMessageOrder(&wg)

	app.Get("/", func(c *fiber.Ctx) {

	})

	app.Listen(":6000")
}
