package auth

import (
	kafka2 "GoWayTaxiAuthService/internal/fetcher/kafka"
	"GoWayTaxiAuthService/pkg/database"
	"GoWayTaxiAuthService/pkg/models"
	"errors"
	"log"
	"strconv"
)

func registerEntity(entity interface{}, role string, req models.AuthRequest) (string, error) {
	jwtKey, err := getJWTKey()
	if err != nil {
		return "", err
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return "", err
	}

	switch e := entity.(type) {
	case *models.User:
		e.Role = role
		e.Email = req.Email
		e.Name = req.Name
		e.Password = hashedPassword

		result := database.DB.Create(e)
		if result.Error != nil {
			return "", result.Error
		}

		InitKafka()
		kafka2.SendMessage("user-created", e)
		kafka2.SendMessage("user-get", e)
		return generateToken(strconv.Itoa(e.Id), e.Email, e.Role, jwtKey)

	case *models.Driver:
		e.Role = role
		e.Email = req.Email
		e.Name = req.Name
		e.Password = hashedPassword

		result := database.DB.Create(e)
		if result.Error != nil {
			return "", result.Error
		}

		InitKafka()
		kafka2.SendMessage("user-created", e)
		kafka2.SendMessage("user-get", e)
		return generateToken(strconv.Itoa(e.Id), e.Email, e.Role, jwtKey)
	}

	return "", errors.New("unsupported entity type")
}

func authenticateEntity(entity interface{}, req models.AuthRequest) (string, error) {
	jwtKey, err := getJWTKey()
	if err != nil {
		return "", err
	}

	switch e := entity.(type) {
	case *models.User:
		result := database.DB.Where("email = ?", req.Email).First(e)
		if result.Error != nil {
			return "", result.Error
		}
		if !checkPasswordHash(req.Password, e.Password) {
			return "", errors.New("wrong password")
		}
		return generateToken(strconv.Itoa(e.Id), e.Email, e.Role, jwtKey)

	case *models.Driver:
		result := database.DB.Where("email = ?", req.Email).First(e)
		if result.Error != nil {
			return "", result.Error
		}
		if !checkPasswordHash(req.Password, e.Password) {
			return "", errors.New("wrong password")
		}
		return generateToken(strconv.Itoa(e.Id), e.Email, e.Role, jwtKey)
	}

	return "", errors.New("unsupported entity type")
}

func InitKafka() {
	topics := []string{"user-created", "user-get"}
	for _, t := range topics {
		if err := kafka2.CreateTopic("localhost:9092", t); err != nil {
			log.Printf("topic %s already exists or failed: %v", t, err)
		}
	}
}
