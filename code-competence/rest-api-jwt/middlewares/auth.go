package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateToken(userID uint) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
