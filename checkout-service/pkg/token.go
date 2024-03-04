package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	return err == nil && token.Valid
}