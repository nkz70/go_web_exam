package persistence

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"webexam/domain/model"
	"webexam/domain/repository"
)

type questionPersistence struct {
	con *gorm.DB
}

func NewQuestionPersistence(con *gorm.DB) repository.QuestionRepository {
	return &questionPersistence{con}
}

func (qp *questionPersistence) FetchList() (*[]model.Question, error) {
	var questions []model.Question

	c := qp.con

	// c := up.createDBClient(fc)

	if err := c.Find(&questions).Error; err != nil {
		return nil, errors.WithMessage(err, "failed find question")
	}

	return &questions, nil
}

func (qp *questionPersistence) Find(id int64) (*model.Question, error) {
	var question model.Question

	if err := qp.con.First(&question, id).Error; err != nil {
		return nil, errors.WithMessage(err, "failed find question")
	}

	return &question, nil
}

func (qp *questionPersistence) Create(question *model.Question) (*model.Question, error) {
	if err := qp.con.Create(question).Error; err != nil {
		return nil, errors.WithMessage(err, "failed create question")
	}

	return question, nil
}

func (qp *questionPersistence) Update(question *model.Question) (*model.Question, error) {
	if err := qp.con.Save(question).Error; err != nil {
		return nil, errors.WithMessage(err, "failed update question")
	}

	return question, nil
}

func (qp *questionPersistence) Delete(id int64) error {
	var question model.Question

	if err := qp.con.Delete(&question, id).Error; err != nil {
		return errors.WithMessage(err, "failed delete question")
	}

	return nil
}

func (qp *questionPersistence) createDBClient(fc *model.FindClause) *gorm.DB {
	w := qp.con.Order("id desc")
	w = w.Limit(fc.Limit)
	w = w.Where(fc)

	return w
}
