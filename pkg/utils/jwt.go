package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret string

type Claims struct {
	ID    string
	Email string
  IsGuest bool
	jwt.RegisteredClaims
}

func GenerateTokenPair(userID, email string) (map[string]string, error) {
	var secret = os.Getenv("JWT_SECRET")
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		userID,
		email,
    false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}).SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		userID,
		email,
    false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}).SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}

	tokenPair := make(map[string]string)
	tokenPair["accessToken"] = accessToken
	tokenPair["refreshToken"] = refreshToken

	return tokenPair, nil

}
