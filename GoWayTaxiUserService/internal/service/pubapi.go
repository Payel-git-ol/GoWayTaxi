package service

import "GoWayTaxiUserService/pkg/models"

func SaveUser(user models.User) (string, error) {
	return save(user)
}

func SaveDriver(driver models.DriverUs) (string, error) {
	return save(driver)
}
