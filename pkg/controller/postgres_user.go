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

func (r *BookstorePostgres) CreateUser(cnt context.Context, user domain.User) bool {
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	r.Db.Create(&domain.User{
		ID:           id,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		RegisteredAt: user.RegisteredAt,
	})
	return true
}
func (r *BookstorePostgres) SignInUser(cnt context.Context, email, password string) (domain.User, bool) {
	var user domain.User
	r.Db.Where("email = ? AND password = ?", email, password).Find(&user)
	return user, true
}
