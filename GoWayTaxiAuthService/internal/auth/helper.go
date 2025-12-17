package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func getJWTKey() ([]byte, error) {
	jwtKey := []byte(os.Getenv("JWT_TOKEN_AU"))
	if len(jwtKey) == 0 {
		return nil, errors.New("jwt token is empty")
	}
	return jwtKey, nil
}

func generateToken(id string, email string, role string, jwtKey []byte) (string, error) {
	var people string

	if role == "driver" {
		people = "driver"
	} else {
		people = "user"
	}

	claims := jwt.MapClaims{
		people + "_id": id,
		"email":        email,
		"Role":         role,
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
