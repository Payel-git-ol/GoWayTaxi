package main

import (
	"GoWayTaxiPricingService/internal/fetcher/kafka"
	"GoWayTaxiPricingService/metrics"
	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
)

func main() {
	app := fiber.New()
	var wg sync.WaitGroup
	wg.Add(1)
	go kafka.GetMessageOrder(&wg)

	app.Get("/", func(c fiber.Ctx) {})

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		log.Println("metrics server on :9100")
		if err := http.ListenAndServe(":9100", mux); err != nil {
			log.Fatalf("metrics server error: %v", err)
		}
	}()

	metrics.Init()
	app.Listen(":6000")
}
