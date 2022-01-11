package repository

import (
	"golang-template-module/internal/models"
	db "golang-template-module/pkg/database"
)

func GetById(id uint64) (*db.User, error) {
	user := db.User{}
	err := db.DB.First(&user, id).Error

	if err != nil {
		return nil, models.NotFoundError
	}

	return &user, nil
}

func GetByEmail(email string) (*db.User, error) {
	user := db.User{}
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, models.NotFoundError
	}
	return &user, nil
}

func InsertUser(m *models.UserAuthModel) (*db.User, error) {
	user := db.User{Email: m.Email, PasswordHash: m.Password}
	err := db.DB.Create(&user).Error
	if err != nil {
		return nil, models.ExistFoundError
	}
	return &user, nil
}

func GetList() ([]*db.User, error) {
	users := []*db.User{}
	err := db.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
