package common

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(userID, tokenType, secret string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType, // "access" atau "refresh"
		"exp":     time.Now().Add(duration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
