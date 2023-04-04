package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	domain "github.com/Onelvay/docker-compose-project/pkg/domain"
	jwt "github.com/golang-jwt/jwt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}
type UserDbActioner interface {
	CreateUser(cnt context.Context, user domain.User) bool
	SignInUser(context.Context, string, string) (domain.User, bool)
}
type UserController struct {
	repo UserDbActioner
	// hasher PasswordHasher

	hmacSecret []byte
}

func NewUserController(db UserDbActioner) *UserController {
	return &UserController{repo: db}
}
func (s *UserController) SignUp(ctx context.Context, inp domain.SignUpInput) bool {
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
	return s.repo.CreateUser(ctx, user)
}
func (s *UserController) SignIn(ctx context.Context, inp domain.SignInInput) (string, error) {
	user, _ := s.repo.SignInUser(ctx, inp.Email, inp.Password)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   user.ID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	})
	return token.SignedString(s.hmacSecret)
}
func (s *UserController) ParseToken(ctx context.Context, token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("aaa")
		}
		return s.hmacSecret, nil
	})
	if err != nil {
		return "", err
	}
	if !t.Valid {
		return "", errors.New("invalid token")
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}
	subject, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid subject")
	}
	return subject, nil
}
