package model

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type Question struct {
	ID              int64            `json:"id"`
	Content         string           `json:"content"`
	Image           string           `json:"image"`
	Number          int64            `json:"number"`
	ShortName       string           `json:"short_name"`
	FormType        uint             `json:"form_type"`
	Status          int64            `json:"status"`
	QuestionAnswers []QuestionAnswer `json:"question_answers"`
	DeletedAt       null.Time        `json:"deleted_at"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}
