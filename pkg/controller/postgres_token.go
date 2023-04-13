package controller

import (
	"context"
	"errors"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"gorm.io/gorm"
)

type TokenPostgres struct {
	Db *gorm.DB
}

func NewTokenDbController(db *gorm.DB) *TokenPostgres {
	return &TokenPostgres{Db: db}
}

func (r *TokenPostgres) CreateToken(cnt context.Context, token domain.Refresh_token) bool {
	r.Db.Create(&token)
	return true
}
func (r *TokenPostgres) GetToken(cxt context.Context, token string) domain.Refresh_token {
	var d domain.Refresh_token
	r.Db.Where("token= ?", token).Find(&d)
	r.Db.Delete(&d)
	return d
}
func (r *TokenPostgres) GetUserIdByToken(token string) (string, error) {
	var d domain.Refresh_token
	if token == "" {
		return "", errors.New("token is empty")
	}
	res := r.Db.Where("token= ?", token).Find(&d)
	if res.RowsAffected == 0 {
		return "", errors.New("user not found")
	}
	return d.UserId, nil
}
