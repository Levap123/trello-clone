package jwt

import (
	"os"
	"time"

	errs "github.com/Levap123/trello-clone/pkg/errors"
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
		return "", err
	}
	return tokenString, nil
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrInvalidSign
		}
		return sign, nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errs.ErrInvalidClaims
	}
	return claims.UserId, nil
}
