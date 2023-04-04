package service

import (
	"context"
	"time"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}
type UserController interface {
	Create(cnt context.Context, user domain.User) bool
}
type Users struct {
	repo UserController
	// hasher PasswordHasher

	// hmacSecret []byte
}

func NewUsers(db UserController) *Users {
	return &Users{repo: db}
}
func (s *Users) SignUp(ctx context.Context, inp domain.SignUpInput) bool {
	// password, err := s.hasher.Hash(inp.Password)
	// if err != nil {
	// 	return err
	// }
	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     inp.Password,
		RegisteredAt: time.Now(),
	}
	return s.repo.Create(ctx, user)
}
