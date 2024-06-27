package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"

	"webexam/domain/model"
)

type UserValidator interface {
	Validate(user *model.User) error
}

type userValidator struct{}

func NewUserValidator() UserValidator {
	return &userValidator{}
}

func (uv *userValidator) Validate(user *model.User) error {
	return validation.ValidateStruct(user,
		validation.Field(
			&user.FirstName,
			validation.Required.Error("first name is required"),
			validation.Length(1, 20).Error("limited min 1 max 20 char"),
		),
		validation.Field(&user.LastName, validation.Required, validation.Length(1, 20)),
		validation.Field(&user.Age, validation.Min(1), validation.Max(100)),
		validation.Field(&user.Gender, validation.Min(1), validation.Max(2)),
	)
}
