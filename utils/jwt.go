package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func init() {
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET is not set")
	}
	log.Println("JWT_SECRET loaded successfully")
}

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := claims["userID"].(float64)
		if !ok {
			return 0, errors.New("invalid token claims")
		}
		return uint(userIDFloat), nil
	}
	return 0, errors.New("invalid token")
}
