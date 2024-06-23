package model

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type User struct {
	ID           int64        `json:"id"`
	LastName     string       `json:"last_name"`
	FirstName    string       `json:"first_name"`
	Age          null.Int64   `json:"age"`
	Gender       null.Int64   `json:"gender"`
	Score        null.Int64   `json:"score"`
	Time         null.Int64   `json:"time"`
	Status       null.Int64   `json:"status"`
	UserAnswers  []UserAnswer `json:"user_answers"`
	TestDatetime null.Time    `json:"test_datetime"`
	DeletedAt    null.Time    `json:"deleted_at"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type FindClause struct {
	LastName  string
	FirstName string
	Limit     int
}
