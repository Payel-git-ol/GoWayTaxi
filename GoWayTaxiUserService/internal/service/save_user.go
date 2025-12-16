package service

import (
	"GoWayTaxiUserService/pkg/database"
	"GoWayTaxiUserService/pkg/models"
	"fmt"
	"log"
)

func save(entity interface{}) (string, error) {
	switch e := entity.(type) {
	case *models.User:
		var existingUser models.User

		result := database.DB.Where("email = ?", e.Email).First(&existingUser)

		if result.Error == nil {
			log.Printf("⚠️ User with email %s already exists (ID: %d)", e.Email, existingUser.Id)
		}

		result = database.DB.Create(e)
		if result.Error != nil {
			log.Printf("❌ Failed to save user %s: %v", e.Email, result.Error)
			return "", result.Error
		}

		return fmt.Sprintf("✅ User saved: %s (ID: %d)", e.Email, e.Id), nil

	case *models.DriverUs:
		var existingDriver models.DriverUs

		result := database.DB.Where("email = ?", e.Email).First(&existingDriver)

		if result.Error == nil {
			log.Printf("⚠️ User with email %s already exists (ID: %d)", e.Email, existingDriver.Id)
		}

		if result.Error == nil {
			log.Printf("⚠️ User with email %s already exists (ID: %d)", e.Email, existingDriver.Id)
		}

		result = database.DB.Create(e)
		if result.Error != nil {
			log.Printf("❌ Failed to save user %s: %v", e.Email, result.Error)
			return "", result.Error
		}

		return fmt.Sprintf("✅ Driver saved: %s (ID: %d)", e.Email, e.Id), nil
	}

	return "✅Saved", nil
}
