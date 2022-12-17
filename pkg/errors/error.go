package errs

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrUnique            = errors.New("user with this name or username is already exists")
	ErrInvalidEmail      = errors.New("user with this email does not exist")
	ErrPasswordIncorrect = errors.New("password is incorrect")
	ErrInvalidSign       = errors.New("invalid signing method")
	ErrInvalidClaims     = errors.New("token claims type invalid")
)

func Fail(err error, place string) error {
	return fmt.Errorf("%s: %s", place, err.Error())
}

func WebFail(status int) error {
	return fmt.Errorf(http.StatusText(status))
}
