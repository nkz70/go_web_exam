package repository

import (
	"webexam/domain/model"
)

type QuestionRepository interface {
	FetchList() (*[]model.Question, error)
	Find(id int64) (*model.Question, error)
	Create(user *model.Question) (*model.Question, error)
	Update(user *model.Question) (*model.Question, error)
	Delete(id int64) error
}
