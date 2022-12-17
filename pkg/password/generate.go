package pass

import "golang.org/x/crypto/bcrypt"

func GeneratePasswordHash(password string) (string, error) {
	buffer, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(buffer), err
}

func ComparePassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
