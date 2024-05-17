package configs

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET = []byte("secret")

func GenerateJWT(email string, name string, id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["name"] = name
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	tokenString, err := token.SignedString(SECRET)

	if err != nil {
		log.Fatal("Unable to generate JWT")
		return "", err
	}

	return tokenString, nil
}