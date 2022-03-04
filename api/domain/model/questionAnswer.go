package model

import "gorm.io/gorm"

type QuestionAnswer struct {
	gorm.Model
	FormType      string
	QuestionID    uint
	Answer        string
	LabelPosition uint
	Label         string
}
