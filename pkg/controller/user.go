package controller

import (
	"context"
	"time"

	domain "github.com/Onelvay/docker-compose-project/pkg/domain"
	i "github.com/Onelvay/docker-compose-project/pkg/service"
)

type UserController struct {
	userRepo  i.UserDbActioner
	tokenRepo i.TokenDbActioner
	hasher    i.PasswordHasher
	orderRepo i.Transactioner

	hmacSecret []byte
}

func NewUserController(db i.UserDbActioner, tdb i.TokenDbActioner, hash i.PasswordHasher, or i.Transactioner) UserController {
	return UserController{userRepo: db, tokenRepo: tdb, hasher: hash, orderRepo: or}
}
func (s *UserController) SignUp(ctx context.Context, inp domain.SignUpInput) error {
	password := s.hasher.Hash(inp.Password) //хешируем пароль
	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}
	return s.userRepo.CreateUser(ctx, user)
}
func (s *UserController) SignIn(ctx context.Context, inp domain.SignInInput) (string, string, error) {
	password := s.hasher.Hash(inp.Password)
	user, _ := s.userRepo.SignInUser(ctx, inp.Email, password)

	return s.GenerateTokens(ctx, user.ID)
}
