package errs

import (
	"errors"
	"fmt"
	"net/http"
)

var UniqueErr = errors.New("user with this name or username is already exists")

func Fail(err error, place string) error {
	return fmt.Errorf("%s: %s", place, err.Error())
}

func WebFail(status int) error {
	return fmt.Errorf(http.StatusText(status))
}
