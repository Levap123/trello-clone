package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var sign = os.Getenv("SIGN")

func GenerateJwt(id int) (string, error) {
	fmt.Println(sign)
	claims := jwt.MapClaims{
		"exp": time.Now().Add(72 * time.Hour),
		"id":  id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(sign)
}
