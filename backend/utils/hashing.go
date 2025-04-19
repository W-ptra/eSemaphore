package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(plainPassword string) (string, error) {
	saltRound := 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), saltRound)
	return string(bytes), err
}

func ComparePassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
