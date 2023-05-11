package encrypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password, hashPwd string) (bool, error) {
	isValid := false
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashPwd)); err != nil {
		err = errors.New("password is incorrect")
		return isValid, err
	}

	isValid = true
	return isValid, nil
}
