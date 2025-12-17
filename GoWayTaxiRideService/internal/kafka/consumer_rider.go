package kafka

import (
	"RideService/internal/service"
	modelsrider "RideService/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"sync"
)

func GetMessageUserAndDriver(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
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

		if err := processMessage(m.Value); err != nil {
			fmt.Println(err)
		}
	}
}

func processMessage(data []byte) error {
	fmt.Printf("Consumer started")
	var base struct {
		Role string `json:"role"`
	}
	if err := json.Unmarshal(data, &base); err != nil {
		return err
	}

	switch base.Role {
	case "user":
		var user modelsrider.User
		if err := json.Unmarshal(data, &user); err != nil {
			return err
		}
		_, err := service.Save(&user)
		return err

	case "driver":
		var driver modelsrider.Driver
		if err := json.Unmarshal(data, &driver); err != nil {
			return err
		}
		_, err := service.Save(&driver)
		return err

	default:
		return fmt.Errorf("unknown role: %s", base.Role)
	}
}
