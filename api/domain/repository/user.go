package repository

import (
	"webxam/domain/model"
)

type UserRepository interface {
	FetchList() (*[]model.User, error)
	Find(id int64) (*model.User, error)
}
