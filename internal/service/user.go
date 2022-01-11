package service

import (
	"golang-template-module/internal/models"
	"golang-template-module/internal/repository"
)

func GetUserService(id uint64) (model *models.UserBaseModel, err error) {
	user, err := repository.GetById(id)
	if err != nil {
		return nil, err
	}
	model = new(models.UserBaseModel)
	model.Email = user.Email
	model.Id = user.Id
	return model, nil
}

func GetUserListService() (*models.UserListBaseModel, error) {
	user, err := repository.GetList()
	if err != nil {
		return nil, err
	}
	ListModels := new(models.UserListBaseModel)
	for _, usr := range user {
		ListModels.Users = append(ListModels.Users, &models.UserBaseModel{Email: usr.Email, Id: usr.Id})
	}
	return ListModels, nil
}
