package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
