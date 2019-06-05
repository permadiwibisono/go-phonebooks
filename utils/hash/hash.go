package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(str string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func Compare(hashed string, plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
	if err != nil {
		return false, err
	}
	return true, nil
}
