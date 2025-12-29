package kafka

import (
	"RideService/internal/core/service/message"
	"RideService/metrics"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"sync"
)

func GetMessageUserAndDriver(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "user-get",
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

		metrics.KafkaMessagesIn.Inc()

		if err := message.ProcessMessageAuth(m.Value); err != nil {
			fmt.Println(err)
		}
	}
}
