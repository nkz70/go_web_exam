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
	FindUser(ur *UserRequest) (*[]model.User, error)
}

type userController struct {
	UserRepository repository.UserRepository
}

func NewUserController(ur repository.UserRepository) UserController {
	return &userController{
		UserRepository: ur,
	}
}

func (uc *userController) FindUser(ur *UserRequest) (*[]model.User, error) {
	// nur := validate(ur)
	users, err := uc.UserRepository.Find()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// func validate(ur UserRequest) {

// }
