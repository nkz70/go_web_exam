package model

import (
	"time"
)

type User struct {
	ID           string       `json:"id"`
	LastName     string       `json:"last_name"`
	FirstName    string       `json:"first_name"`
	Age          int64        `json:"age"`
	Gender       int64        `json:"gender"`
	Score        int64        `json:"score"`
	Time         int64        `json:"time"`
	Status       int64        `json:"status"`
	UserAnswers  []UserAnswer `json:"user_answers"`
	TestDatetime time.Time    `json:"test_datetime"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    time.Time    `json:"deleted_at"`
}

type FindClause struct {
	LastName  string
	FirstName string
	Limit     int
}
