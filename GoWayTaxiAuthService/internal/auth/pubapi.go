package auth

import (
	"GoWayTaxiAuthService/pkg/models"
	"GoWayTaxiAuthService/pkg/models/request"
)

func RegisterUser(req request.AuthRequest) (string, error) {
	return registerEntity(&models.User{}, "user", req)
}

func RegisterDriver(req request.AuthRequest) (string, error) {
	return registerEntity(&models.Driver{}, "driver", req)
}

func AuthenticateUser(req request.AuthRequest) (string, error) {
	return authenticateEntity(&models.User{}, req)
}

func AuthenticateDriver(req request.AuthRequest) (string, error) {
	return authenticateEntity(&models.Driver{}, req)
}
