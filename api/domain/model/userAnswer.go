package model

import "gorm.io/gorm"

type UserAnswer struct {
	gorm.Model
	Answer     string
	Correct    int64
	UserID     uint
	QuestionID uint
}
