package utils

import (
	"cafekalaa/api/app/model"
	"log"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("best_secret_key")

func MakeRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {

	hmacSecret := []byte(jwtKey)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {

		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func ExtractClaimsForRefresh(tokenStr string) (jwt.MapClaims, bool) {

	hmacSecret := []byte(jwtKey)
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func MakeTokenFromUUID(uuid uuid.UUID) string {

	expirationTime := time.Now().Add(15 * time.Second)

	claims := &model.Claims{
		ID: uuid,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		panic(err)
	}

	return tokenString
}
