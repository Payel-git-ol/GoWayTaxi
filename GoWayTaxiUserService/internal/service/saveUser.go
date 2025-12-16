package service

import (
	"GoWayTaxiUserService/pkg/database"
	"GoWayTaxiUserService/pkg/models"
	"log"
)

func SaveUser(user models.User) error {
	if database.DB == nil {
		log.Println("❌ Database connection is not initialized")
		return nil
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)

	if result.Error == nil {
		log.Printf("⚠️ User with email %s already exists (ID: %d)",
			user.Email, existingUser.Id)
		return nil
	}

	result = database.DB.Create(&user)
	if result.Error != nil {
		log.Printf("❌ Failed to save user %s: %v", user.Email, result.Error)
		return result.Error
	}

	log.Printf("✅ User saved: %s (ID: %d)", user.Email, user.Id)
	return nil
}

func SaveDriver(driver models.DriverUs) error {
	if database.DB == nil {
		log.Println("❌ Database connection is not initialized")
		return nil
	}

	var existingDriver models.DriverUs
	result := database.DB.Where("email = ?", driver.Email).First(&existingDriver)

	if result.Error == nil {
		log.Printf("⚠️ Driver with email %s already exists (ID: %d)",
			driver.Email, existingDriver.Id)
		return nil
	}

	result = database.DB.Create(&driver)
	if result.Error != nil {
		log.Printf("❌ Failed to save driver %s: %v", driver.Email, result.Error)
		return result.Error
	}

	log.Printf("✅ Driver saved: %s (ID: %d)", driver.Email, driver.Id)
	return nil
}
