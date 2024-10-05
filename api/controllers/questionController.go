package controllers

import (
	"webxam/domain/model"
	"webxam/domain/repository"

	"go.uber.org/zap/zapcore"
)

type QuestionRequest struct {
	Content   string `json:"content"`
	Image     string `json:"image"`
	Number    int64  `json:"number"`
	ShortName string `json:"short_name"`
	Status    int64  `json:"status"`
}

func (qr QuestionRequest) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("content", qr.Content)
	enc.AddString("image", qr.Image)
	enc.AddInt64("number", qr.Number)
	enc.AddString("short_name", qr.ShortName)
	enc.AddInt64("status", qr.Status)
	return nil
}

type QuestionController interface {
	FetchQuestionList(qr *QuestionRequest) (*[]model.Question, error)
	FindQuestion(id int64) (*model.Question, error)
	CreateQuestion(qr *QuestionRequest) (*model.Question, error)
	UpdateQuestion(id int64, ur *QuestionRequest) (*model.Question, error)
	DeleteQuestion(id int64) error
}

type questionController struct {
	QuestionRepository repository.QuestionRepository
}

func NewQuestionController(qr repository.QuestionRepository) QuestionController {
	return &questionController{
		QuestionRepository: qr,
	}
}

func (qc *questionController) FetchQuestionList(qr *QuestionRequest) (*[]model.Question, error) {
	questions, err := qc.QuestionRepository.FetchList()
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (qc *questionController) FindQuestion(id int64) (*model.Question, error) {
	question, err := qc.QuestionRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (qc *questionController) CreateQuestion(qr *QuestionRequest) (*model.Question, error) {
	question := model.Question{
		Content:   qr.Content,
		Image:     qr.Image,
		Number:    qr.Number,
		ShortName: qr.ShortName,
		Status:    qr.Status,
	}

	res, err := qc.QuestionRepository.Create(&question)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (qc *questionController) UpdateQuestion(id int64, qr *QuestionRequest) (*model.Question, error) {
	question := model.Question{
		Content:   qr.Content,
		Image:     qr.Image,
		Number:    qr.Number,
		ShortName: qr.ShortName,
		Status:    qr.Status,
	}

	res, err := qc.QuestionRepository.Update(&question)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (qc *questionController) DeleteQuestion(id int64) error {
	if err := qc.QuestionRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

// func validate(ur UserRequest) {

// }
