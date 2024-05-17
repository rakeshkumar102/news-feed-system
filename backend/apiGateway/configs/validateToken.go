package configs

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET = []byte("secret")

func ValidateToken(token string) (jwt.MapClaims, bool) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error parsing token")
		}
		return SECRET, nil
	})

	if err != nil {
		return nil, false
	}

	if !tk.Valid {
		return nil, false
	}

	claims, ok := tk.Claims.(jwt.MapClaims)

	if !ok {
		return nil, false
	}

	return claims, true
}