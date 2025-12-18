package kafka

import (
	"RideService/metrics"
	"RideService/pkg/models/request"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func SendMessagePricing(req request.RequestOrder, topic string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{Value: jsonData},
	)

	metrics.KafkaMessagesOut.Inc()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Отправлено в топик '%s': %v\n", topic, string(jsonData))
}
