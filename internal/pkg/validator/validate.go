package validator

import (
	"regexp"
	"unicode"
)

var (
	regexEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func IsPasswordValid(password string) bool {
	var hasUpper, hasLower bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		}
	}

	return len(password) >= 8 && hasUpper && hasLower
}

func IsPasswordConfirmed(password string, passwordConfirmation string) bool {
	return password == passwordConfirmation
}

func IsValidEmail(email string) bool {
	return regexEmail.MatchString(email)
}
