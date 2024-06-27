package controllers

import (
	"go.uber.org/zap/zapcore"

	"webexam/domain/model"
	"webexam/domain/repository"
)

type QuestionRequest struct {
	Content         string          `json:"content"`
	Image           string          `json:"image"`
	Number          int64           `json:"number"`
	ShortName       string          `json:"short_name"`
	FormType        uint            `json:"form_type"`
	Status          int64           `json:"status"`
	QuestionAnswers QuestionAnswers `json:"question_answers"`
}

func (qr QuestionRequest) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("content", qr.Content)
	enc.AddString("image", qr.Image)
	enc.AddInt64("number", qr.Number)
	enc.AddString("short_name", qr.ShortName)
	enc.AddUint("form_type", qr.FormType)
	enc.AddInt64("status", qr.Status)
	enc.AddArray("question_answers", qr.QuestionAnswers)
	return nil
}

type QuestionAnswers []QuestionAnswer

func (qas QuestionAnswers) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, qa := range qas {
		if err := enc.AppendObject(qa); err != nil {
			return err
		}
	}

	return nil
}

type QuestionAnswer struct {
	Answer        string `json:"answer"`
	LabelPosition uint   `json:"label_position"`
	Label         string `json:"label"`
}

func (qa QuestionAnswer) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("answer", qa.Answer)
	enc.AddUint("label_position", qa.LabelPosition)
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
	qs := make([]model.QuestionAnswer, len(qr.QuestionAnswers))
	for i, v := range qr.QuestionAnswers {
		qs[i] = model.QuestionAnswer{
			Answer:        v.Answer,
			LabelPosition: v.LabelPosition,
			Label:         v.Label,
		}
	}

	question := model.Question{
		Content:         qr.Content,
		Image:           qr.Image,
		Number:          qr.Number,
		ShortName:       qr.ShortName,
		FormType:        qr.FormType,
		Status:          qr.Status,
		QuestionAnswers: qs,
	}

	res, err := qc.QuestionRepository.Create(&question)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (qc *questionController) UpdateQuestion(id int64, qr *QuestionRequest) (*model.Question, error) {
	question, err := qc.QuestionRepository.Find(id)
	if err != nil {
		return nil, err
	}

	updateQuestion := createUpdateQuestion(question, qr)
	res, err := qc.QuestionRepository.Update(updateQuestion)
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

func createUpdateQuestion(question *model.Question, qr *QuestionRequest) *model.Question {
	question.Content = qr.Content
	question.Image = qr.Image
	question.Number = qr.Number
	question.ShortName = qr.ShortName
	question.Status = qr.Status
	return question
}
