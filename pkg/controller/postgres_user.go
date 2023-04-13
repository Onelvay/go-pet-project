package controller

import (
	"context"
	"strings"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPostgres struct {
	Db *gorm.DB
}

func NewUserDbController(db *gorm.DB) *BookstorePostgres {
	return &BookstorePostgres{Db: db}
}

func (r *BookstorePostgres) CreateUser(cnt context.Context, user domain.User) error {
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	res := r.Db.Create(&domain.User{
		ID:           id,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		RegisteredAt: user.RegisteredAt,
	})
	return res.Error
}
func (r *BookstorePostgres) SignInUser(cnt context.Context, email, password string) (domain.User, error) {
	var user domain.User
	res := r.Db.Where("email = ? AND password = ?", email, password).Find(&user)
	return user, res.Error
}
