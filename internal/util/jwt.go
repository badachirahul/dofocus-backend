package util

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}
