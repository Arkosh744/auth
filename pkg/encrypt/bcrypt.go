package encrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}
