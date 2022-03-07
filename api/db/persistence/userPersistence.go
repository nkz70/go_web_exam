package persistence

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"webxam/domain/model"
	"webxam/domain/repository"
)

type userPersistence struct {
	con *gorm.DB
}

func NewUserPersistence(con *gorm.DB) repository.UserRepository {
	return &userPersistence{con}
}

func (up *userPersistence) FetchList() (*[]model.User, error) {
	var users []model.User

	c := up.con

	// c := up.createDBClient(fc)

	if err := c.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (up *userPersistence) Find(id int64) (*model.User, error) {
	var user model.User

	if err := up.con.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (up *userPersistence) Create(user *model.User) (*model.User, error) {
	if err := up.con.Select("FirstName", "LastName", "Age", "Gender").Create(user).Error; err != nil {
		return nil, errors.WithMessage(err, "failed create user")
	}

	return user, nil
}

func (up *userPersistence) Delete(id int64) error {
	var user model.User

	if err := up.con.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

func (up *userPersistence) createDBClient(fc *model.FindClause) *gorm.DB {
	w := up.con.Order("id desc")
	w = w.Limit(fc.Limit)
	w = w.Where(fc)

	return w
}
