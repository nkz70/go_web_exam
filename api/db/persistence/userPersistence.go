package persistence

import (
	"gorm.io/gorm"

	"webxam/domain/model"
	"webxam/domain/repository"
)

type WhereClause struct {
	LastName  string
	FirstName string
}

type userPersistence struct {
	con *gorm.DB
}

func NewUserPersistence(con *gorm.DB) repository.UserRepository {
	return &userPersistence{con}
}

func (up *userPersistence) Find() (*[]model.User, error) {
	var users []model.User

	w := up.con.Order("id desc")
	w = w.Limit(20)
	wc := WhereClause{FirstName: "", LastName: ""}
	w = w.Where(&wc)

	if err := w.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}
