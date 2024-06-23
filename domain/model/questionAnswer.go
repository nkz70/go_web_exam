package model

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type QuestionAnswer struct {
	ID            int64 `json:"id"`
	QuestionID    uint
	Answer        string    `json:"answer"`
	LabelPosition uint      `json:"label_position"`
	Label         string    `json:"label"`
	DeletedAt     null.Time `json:"deleted_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
