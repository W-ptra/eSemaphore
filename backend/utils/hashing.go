package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(plainPassword string) string {
	saltRound := 10
	bytes, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), saltRound)
	return string(bytes)
}

func ComparePassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
