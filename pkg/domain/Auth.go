package domain

import (
	"time"

	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Refresh_token struct {
	UserId    string `json:"userId"`
	Token     string `json:"token" gorm:"primaryKey"`
	ExpiresAt time.Time

	User User `gorm:"references:id"`
}

type (
	SignUpInput struct {
		Name     string `json:"name" validate:"required,gte=2"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=4"`
	}

	SignInInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=4"`
	}
)

func (i SignUpInput) Validate() error {
	return validate.Struct(i)
}
func (i SignInInput) Validate() error {
	return validate.Struct(i)
}
