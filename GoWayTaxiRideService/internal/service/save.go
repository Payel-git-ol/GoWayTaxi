package service

import (
	datarider "RideService/pkg/database"
	modelrider "RideService/pkg/models"
	"fmt"
	"log"
)

func Save(entity interface{}) (string, error) {
	switch e := entity.(type) {
	case *modelrider.User:
		var existingUser modelrider.User
		result := datarider.DB.Where("email = ?", e.Email).First(&existingUser)

		if result.RowsAffected > 0 {
			log.Printf("⚠️ User with email %s already exists (ID: %d)", e.Email, existingUser.Id)
			return "", nil
		}

		if err := datarider.DB.Create(e).Error; err != nil {
			return "", err
		}
		return fmt.Sprintf("✅ User saved: %s (ID: %d)", e.Email, e.Id), nil

	case *modelrider.Driver:
		var existingDriver modelrider.Driver
		result := datarider.DB.Where("email = ?", e.Email).First(&existingDriver)

		if result.RowsAffected > 0 {
			log.Printf("⚠️ Driver with email %s already exists (ID: %d)", e.Email, existingDriver.Id)
			return "", nil
		}

		if err := datarider.DB.Create(e).Error; err != nil {
			return "", err
		}
		return fmt.Sprintf("✅ Driver saved: %s (ID: %d)", e.Email, e.Id), nil
	}

	return "", fmt.Errorf("unsupported entity type")
}
