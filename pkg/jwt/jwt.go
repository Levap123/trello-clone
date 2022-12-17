package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var sign = []byte(os.Getenv("SIGN"))

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GenerateJwt(id int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
		},
		id,
	})
	tokenString, err := token.SignedString(sign)
	if err != nil {
		fmt.Println(123)
		return "", err
	}
	return tokenString, nil
}
