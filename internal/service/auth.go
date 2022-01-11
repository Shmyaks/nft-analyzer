package service

import (
	"fmt"
	"golang-template-module/internal/models"
	"golang-template-module/internal/repository"
	hashing "golang-template-module/pkg/hash"
	"golang-template-module/pkg/jwt"
)

func Authorization(m *models.UserAuthModel) (*models.UserTokenModel, error) {
	user, err := repository.GetByEmail(m.Email)
	if err != nil {
		return nil, err
	}

	err = hashing.ComparePassword(user.PasswordHash, m.Password)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Create(user)
	if err != nil {
		return nil, err
	}

	return &models.UserTokenModel{Token: token}, nil
}

func CreateUser(m *models.UserAuthModel) (*models.UserTokenModel, error) {
	_, err := repository.GetByEmail(m.Email)
	fmt.Printf("1")
	if err == nil {
		return nil, models.ExistFoundError
	}
	m.Password = string(hashing.HashPassword(m.Password))
	user, err := repository.InsertUser(m)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Create(user)
	if err != nil {
		return nil, err
	}

	return &models.UserTokenModel{Token: token}, nil
}
