package encrypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(hash, password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	if hash == "" {
		return errors.New("hash cannot be empty")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("passwords do not match")
	}

	return nil
}
