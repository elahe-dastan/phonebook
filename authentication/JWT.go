package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(email string) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = email
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("parham"))

	if err != nil {
		return "", err
	}

	return token, nil
}

//nolint: gofumpt
func ValidateToken(token string) (in bool, i string) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("parham"), nil
	})

	if err != nil {
		return false, ""
	}

	auth := claims["authorized"].(bool)
	//exp := claims["exp"].(float64)
	email := claims["user_id"].(string)

	if auth {
		return true, email
	}

	return false, email
}
