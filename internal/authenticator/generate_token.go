package authenticator

import (
	"time"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error) {
	var user models.User
	claims := &models.Claims{UserId: user.ID.String(), RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
