package encrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password, hashPwd string) bool {
	isValid := false
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashPwd)); err != nil {
		return isValid
	}

	isValid = true
	return isValid
}
