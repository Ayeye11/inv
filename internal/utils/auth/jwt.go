package auth

import (
	"strconv"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(user *models.User, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  strconv.Itoa(user.ID),
		"role": user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
