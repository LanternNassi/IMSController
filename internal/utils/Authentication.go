package utils

import (
	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	"time"

	"fmt"
)

var secretKey = []byte("Nassim")

func GenerateToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func HarshPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidPassword(password string, harshedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(harshedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
