// Package utils contains all the necessary functions
package utils

import "golang.org/x/crypto/bcrypt"

func HashString(a string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(a), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckHash(hash string, a string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(a))
}
