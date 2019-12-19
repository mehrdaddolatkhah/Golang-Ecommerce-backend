package model

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
)

type Claims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}
