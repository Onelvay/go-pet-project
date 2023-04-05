package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
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
	CreateToken(cnt context.Context, token domain.Refresh_token) bool
	GetToken(cxt context.Context, token string) domain.Refresh_token
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
func (s *UserController) SignIn(ctx context.Context, inp domain.SignInInput) (string, string, error) {
	user, _ := s.repo.SignInUser(ctx, inp.Email, inp.Password)

	return s.generateTokens(ctx, user.ID)
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
func (s *UserController) RefreshTokens(ctx context.Context, refreshToken string) (string, string, error) {
	session := s.repo.GetToken(ctx, refreshToken)
	if session.ExpiresAt.Unix() < time.Now().Unix() {
		return "", "", nil
	}
	return s.generateTokens(ctx, session.UserId)
}
func (s *UserController) generateTokens(ctx context.Context, userId string) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	})
	accessToken, err := t.SignedString(s.hmacSecret)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}
	if ok := s.repo.CreateToken(ctx, domain.Refresh_token{
		UserId:    userId,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}); !ok {
		return "", "", nil
	}
	return accessToken, refreshToken, nil
}
func newRefreshToken() (string, error) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
