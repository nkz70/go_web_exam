package controllers

import (
	"webxam/domain/model"
	"webxam/domain/repository"

	"github.com/volatiletech/null/v9"
)

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
	Gender    int64  `json:"gender"`
}

type UserController interface {
	FetchUserList(ur *UserRequest) (*[]model.User, error)
	FindUser(id int64) (*model.User, error)
	CreateUser(ur *UserRequest) (*model.User, error)
	DeleteUser(id int64) error
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

func (uc *userController) CreateUser(ur *UserRequest) (*model.User, error) {
	user := model.User{
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Gender:    null.Int64From(ur.Gender),
		Age:       null.Int64From(ur.Age),
	}

	res, err := uc.UserRepository.Create(&user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *userController) DeleteUser(id int64) error {
	if err := uc.UserRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

// func validate(ur UserRequest) {

// }
