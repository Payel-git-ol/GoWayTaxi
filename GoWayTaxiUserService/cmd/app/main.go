package main

import (
	"GoWayTaxiUserService/internal/kafka"
	"GoWayTaxiUserService/pkg/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"sync"
)

func main() {
	database.InitDB()

	var wg sync.WaitGroup
	wg.Add(1)
	go kafka.GetMessage(&wg)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Authorization", "Content-Type", "User"},
	}))

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":4000")
}
