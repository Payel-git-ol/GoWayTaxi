package kafka

import (
	"GoWayTaxiUserService/internal/service"
	"GoWayTaxiUserService/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"sync"
)

func GetMessage(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user-created",
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Println(string(m.Value))

		rawJSON := string(m.Value)

		fmt.Println("Raw JSON received: %s\n", rawJSON)

		if err := processMessage(m.Value); err != nil {
			fmt.Println(err)
		}
	}
}

func processMessage(data []byte) error {
	fmt.Printf("Consumer started")
	var user models.User
	if err := json.Unmarshal(data, &user); err == nil && user.Role != "" {
		fmt.Printf("✅ Parsed as USER: ID=%d, Name=%s, Email=%s, Role=%s\n",
			user.Id, user.Name, user.Email, user.Role)

		err := service.SaveUser(user)
		if err != nil {
			return err
		}

		return nil
	}

	var driver models.DriverUs
	if err := json.Unmarshal(data, &driver); err == nil && driver.Role != "" {
		fmt.Printf("🚗 Parsed as DRIVER: ID=%d, Name=%s, Email=%s, Role=%s\n",
			driver.Id, driver.Name, driver.Email, driver.Role)

		err := service.SaveDriver(driver)
		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("unknown message format: %s", string(data))
}
