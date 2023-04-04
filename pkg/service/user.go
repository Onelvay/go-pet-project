package service

import (
	"context"
	"time"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	jwt "github.com/golang-jwt/jwt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}
type UserController interface {
	Create(cnt context.Context, user domain.User) bool
	SignInUser(context.Context, string, string) (domain.User, bool)
}
type Users struct {
	repo UserController
	// hasher PasswordHasher

	hmacSecret []byte
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
func (s *Users) SignIn(ctx context.Context, inp domain.SignInInput) (string, error) {
	user, _ := s.repo.SignInUser(ctx, inp.Email, inp.Password)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   user.ID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	})
	return token.SignedString(s.hmacSecret)
}
