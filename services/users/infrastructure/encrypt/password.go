package encrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password, hashPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashPwd))
}
