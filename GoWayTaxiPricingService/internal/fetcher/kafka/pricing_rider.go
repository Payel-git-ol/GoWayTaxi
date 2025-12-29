package kafka

import (
	"GoWayTaxiPricingService/metrics"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type PriceMessage struct {
	OrderId int     `json:"order_id"`
	Price   float64 `json:"price"`
}

func SendMessagePriceInRider(orderId int, res float64, topic string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})
	defer w.Close()

	msg := PriceMessage{OrderId: orderId, Price: res}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{Value: jsonData},
	)
	if err != nil {
		panic(err)
	}

	metrics.KafkaMessagesOut.Inc()
	fmt.Printf("Отправлено в топик '%s': %v\n", topic, string(jsonData))
}
