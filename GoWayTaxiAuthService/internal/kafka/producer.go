package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func SendMessage[T any](topic string, data T) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{Value: jsonData},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Отправлено в топик '%s': %v\n", topic, string(jsonData))
}
