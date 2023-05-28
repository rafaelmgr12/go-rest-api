package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateNewAccessToken returns JWT token
func GenerateNewAccessToken() (string, error) {
	// get the JWT secret key from .env file
	secret := GetValue("JWT_SECRET_KEY")

	// get the JWT token expire time from .env file
	minutesCount, _ := strconv.Atoi(GetValue("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// create a JWT claim object
	claims := jwt.MapClaims{}

	// add expiration time for the token
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// create a new JWT token with the JWT claim object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// convert the token in a string format
	t, err := token.SignedString([]byte(secret))

	// if conversion failed, return the error
	if err != nil {
		return "", err
	}

	// return the token
	return t, nil
}
