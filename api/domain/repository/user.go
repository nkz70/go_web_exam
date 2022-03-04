package repository

import "webxam/domain/model"

type UserRepository interface {
	Find() (*[]model.User, error)
}
