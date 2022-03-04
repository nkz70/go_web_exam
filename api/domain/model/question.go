package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Content   string
	Image     string
	Number    int64
	ShortName string
	Status    int64
}
