package main

import (
	"GoWayTaxiAuthService/internal/core/auth"
	"GoWayTaxiAuthService/metrics"
	"GoWayTaxiAuthService/pkg/database"
	"GoWayTaxiAuthService/pkg/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	app.Post("/api/reg", func(c fiber.Ctx) error {
		var req models.AuthRequest

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		_, err := auth.RegisterUser(req)
		if err != nil {
			return c.Status(fiber.StatusHTTPVersionNotSupported).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status":       fiber.StatusOK,
			"Registration": "success",
		})
	})

	app.Post("/api/auth", func(c fiber.Ctx) error {
		var req models.AuthRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request format",
			})
		}

		if req.Email == "" || req.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email and password are required",
			})
		}

		token, err := auth.AuthenticateUser(req)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}

		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Authentication successful",
			"token":   token,
		})
	})

	app.Post("/api/driver/reg", func(c fiber.Ctx) error {
		var req models.AuthRequest

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		_, err := auth.RegisterDriver(req)
		if err != nil {
			return c.Status(fiber.StatusHTTPVersionNotSupported).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status":       fiber.StatusOK,
			"Registration": "success",
		})
	})

	app.Post("/api/driver/auth", func(c fiber.Ctx) error {
		var req models.AuthRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request format",
			})
		}

		if req.Email == "" || req.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email and password are required",
			})
		}

		token, err := auth.AuthenticateDriver(req)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}

		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Authentication successful",
			"token":   token,
		})
	})

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		log.Println("metrics server on :9100")
		if err := http.ListenAndServe(":9100", mux); err != nil {
			log.Fatalf("metrics server error: %v", err)
		}
	}()

	metrics.Init()
	app.Listen(":3000")
}
