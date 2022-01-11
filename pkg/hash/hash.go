package hashing

import (
	"golang-template-module/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash
}

func ComparePassword(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return models.PasswordOrLoginError
	}
	return err
}
