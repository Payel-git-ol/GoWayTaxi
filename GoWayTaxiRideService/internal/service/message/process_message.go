package message

import (
	"RideService/internal/service/saving"
	modelsrider "RideService/pkg/models"
	"encoding/json"
	"fmt"
)

func ProcessMessage(data []byte) error {
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
		_, err := saving.Save(&user)
		return err

	case "driver":
		var driver modelsrider.Driver
		if err := json.Unmarshal(data, &driver); err != nil {
			return err
		}
		_, err := saving.Save(&driver)
		return err

	default:
		return fmt.Errorf("unknown role: %s", base.Role)
	}
}
