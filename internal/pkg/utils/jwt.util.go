package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenrateJWT(userId int64, email string, jwtSecret []byte, jwtExpiry int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Minute * time.Duration(jwtExpiry)).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
