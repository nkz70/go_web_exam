package model

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type UserAnswer struct {
	ID         int64 `json:"id"`
	Answer     string `json:"answer"`
	Correct    int64 `json:"correct"`
	UserID     uint `json:"user_id"`
	QuestionID uint `json:"question_id"`
	DeletedAt  null.Time `json:"deleted_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
