package domain

import (
	"time"

	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registered_at"`
}
type SignUpInput struct {
	Name     string `json:"name" validate:"required,gte=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4"`
}

func (i SignUpInput) Validate() error {
	return validate.Struct(i)
}
