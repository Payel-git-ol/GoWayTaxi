package main

import (
	"RideService/pkg/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database.InitDB()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app": "GoWayTaxiRide",
		})
	})
}
