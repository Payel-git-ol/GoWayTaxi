package main

import (
	"GoWayTaxiUserService/internal/kafka"
	"GoWayTaxiUserService/metrics"
	"GoWayTaxiUserService/pkg/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	database.InitDB()

	var wg sync.WaitGroup
	wg.Add(1)
	go kafka.GetMessageAuth(&wg)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Authorization", "Content-Type", "User"},
	}))

	jwtKey := []byte(os.Getenv("JWT_TOKEN_US"))
	app.Use(func(c fiber.Ctx) error {
		if c.Path() == "/metrics" {
			return c.Next()
		}
		return jwtMiddleware(jwtKey)(c)
	})

	app.Get("/me", func(c fiber.Ctx) error {
		user := c.Locals("user")

		return c.JSON(user)
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
	app.Listen(":4000")
}

func jwtMiddleware(jwtKey []byte) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		c.Locals("user", token.Claims)
		return c.Next()
	}
}
