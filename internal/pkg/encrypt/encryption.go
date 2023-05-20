package encrypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	fmt.Println("password: ", password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println(string(hashedPasswordBytes))
	return string(hashedPasswordBytes), err
}

func VerifyPassword(hashedPassword string, candidatePassword string) bool {
	fmt.Println(hashedPassword, candidatePassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return err == nil
}
