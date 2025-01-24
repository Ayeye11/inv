package auth

import (
	"fmt"
	"strconv"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(user *models.User, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  strconv.Itoa(user.ID),
		"name": user.Name,
		"role": user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
