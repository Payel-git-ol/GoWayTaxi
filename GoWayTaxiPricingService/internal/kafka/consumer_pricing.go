package kafka

import (
	"GoWayTaxiPricingService/internal/service/message"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"sync"
)

func GetMessageOrder(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "pricing-topic",
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())

		if err != nil {
			break
		}

		fmt.Println(string(m.Value))

		rawJSON := string(m.Value)
		fmt.Printf("Raw JSON received: %s\n\n", rawJSON)

		if _, err := message.ProcessMessage(m.Value); err != nil {
		}
	}
}
