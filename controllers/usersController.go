package controllers

import (
	"github.com/volatiletech/null/v9"
	"go.uber.org/zap/zapcore"

	"webexam/domain/model"
	"webexam/domain/repository"
)

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
	Gender    int64  `json:"gender"`
}

func (ur UserRequest) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("first_name", ur.FirstName)
	enc.AddString("last_name", ur.LastName)
	enc.AddInt64("age", ur.Age)
	enc.AddInt64("gender", ur.Gender)
	return nil
}

type UserController interface {
	FetchUserList(ur *UserRequest) (*[]model.User, error)
	FindUser(id int64) (*model.User, error)
	CreateUser(ur *UserRequest) (*model.User, error)
	UpdateUser(id int64, ur *UserRequest) (*model.User, error)
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

func (uc *userController) UpdateUser(id int64, ur *UserRequest) (*model.User, error) {
	user, err := uc.UserRepository.Find(id)
	if err != nil {
		return nil, err
	}

	updateUser := createUpdateUser(user, ur)
	res, err := uc.UserRepository.Update(updateUser)
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

func createUpdateUser(user *model.User, ur *UserRequest) *model.User {
	user.FirstName = ur.FirstName
	user.LastName = ur.LastName
	user.Gender = null.Int64From(ur.Gender)
	user.Age = null.Int64From(ur.Age)
	return user
}
