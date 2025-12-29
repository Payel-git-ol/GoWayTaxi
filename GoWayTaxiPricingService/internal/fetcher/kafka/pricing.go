package kafka

import (
	"GoWayTaxiPricingService/internal/core/message"
	"GoWayTaxiPricingService/metrics"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"sync"
)

func GetMessageOrder(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "pricing-topic",
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Println("Raw JSON:", string(m.Value))

		metrics.KafkaMessagesIn.Inc()

		orderId, result, err := message.ProcessMessage(m.Value)
		if err != nil {
			fmt.Println("Error processing:", err)
			continue
		}

		SendMessagePriceInRider(orderId, result, "pricing-topic-get-price")
	}
}
