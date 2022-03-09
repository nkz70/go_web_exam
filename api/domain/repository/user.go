package repository

import (
	"webxam/domain/model"
)

type UserRepository interface {
	FetchList() (*[]model.User, error)
	Find(id int64) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id int64) error
}
