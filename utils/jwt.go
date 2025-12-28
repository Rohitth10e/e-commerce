package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func init() {
	if len(secretKey) == 0 {
		fmt.Println("[JWT WARN] SECRET_KEY is not set in environment")
	}
}

func Sign(email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"role":  role,
			"exp":   time.Now().Add(24 * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type ClaimsData struct {
	Email string
	Role  string
}

func Verify(tokenString string) (ClaimsData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return ClaimsData{}, err
	}

	if !token.Valid {
		return ClaimsData{}, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return ClaimsData{}, fmt.Errorf("invalid claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return ClaimsData{}, fmt.Errorf("email claim missing")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return ClaimsData{}, fmt.Errorf("role claim missing")
	}

	return ClaimsData{Email: email, Role: role}, nil

}
