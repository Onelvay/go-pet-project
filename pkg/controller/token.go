package controller

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	jwt "github.com/golang-jwt/jwt"
)

func (s *UserController) ParseToken(ctx context.Context, token string) (string, error) { //ВСЕ ЧТО НИЖЕ СВЯЗАНО С  JWT ЛУЧШЕ НЕ ТРОГАТЬ
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
	session, err := s.tokenRepo.GetToken(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}
	if session.ExpiresAt.Unix() < time.Now().Unix() {
		return "", "", nil
	}
	return s.GenerateTokens(ctx, session.UserId)
}

func (s *UserController) GenerateTokens(ctx context.Context, userId string) (string, string, error) {
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
	if ok := s.tokenRepo.CreateToken(ctx, domain.Refresh_token{
		UserId:    userId,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}); ok != nil {
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
