package service

import "GoWayTaxiUserService/pkg/models"

func SaveUser() (string, error) {
	return save(&models.User{})
}

func SaveDriver() (string, error) {
	return save(&models.DriverUs{})
}
