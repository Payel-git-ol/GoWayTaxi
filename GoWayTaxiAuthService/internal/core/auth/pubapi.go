package auth

import (
	"GoWayTaxiAuthService/pkg/models"
)

func RegisterUser(req models.AuthRequest) (string, error) {
	return registerEntity(&models.User{}, "user", req)
}

func RegisterDriver(req models.AuthRequest) (string, error) {
	return registerEntity(&models.Driver{}, "driver", req)
}

func AuthenticateUser(req models.AuthRequest) (string, error) {
	return authenticateEntity(&models.User{}, req)
}

func AuthenticateDriver(req models.AuthRequest) (string, error) {
	return authenticateEntity(&models.Driver{}, req)
}
