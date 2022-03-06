package controllers

import (
	"webxam/domain/model"
	"webxam/domain/repository"
)

type UserRequest struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

type UserController interface {
	FetchUserList(ur *UserRequest) (*[]model.User, error)
	FindUser(id int64) (*model.User, error)
}

type userController struct {
	UserRepository repository.UserRepository
}

func NewUserController(ur repository.UserRepository) UserController {
	return &userController{
		UserRepository: ur,
	}
}

func (uc *userController) FetchUserList(ur *UserRequest) (*[]model.User, error) {
	users, err := uc.UserRepository.FetchList()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *userController) FindUser(id int64) (*model.User, error) {
	user, err := uc.UserRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// func validate(ur UserRequest) {

// }
