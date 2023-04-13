package service

import (
	"context"

	request "github.com/Onelvay/docker-compose-project/payment/APIrequest"
	"github.com/Onelvay/docker-compose-project/pkg/domain"
)

type Transactioner interface {
	CreateOrder(userId string, orderId string) error
	CreateInfoOrder(request.FinalResponse) error
}

type TokenPostgreser interface {
	CreateToken(cnt context.Context, token domain.Refresh_token) bool
	GetToken(cxt context.Context, token string) domain.Refresh_token
	GetUserIdByToken(token string) (string, error)
}
type BookstorePostgreser interface {
	GetBooks(bool) []domain.Book
	GetBookById(string) (domain.Book, bool)
	GetBooksByName(string) ([]domain.Book, bool)
	DeleteBookById(string) bool
	CreateBook(string, float64, string) bool
	UpdateBook(string, string, string, float64) bool
}
type UserPostgresser interface {
	CreateUser(cnt context.Context, user domain.User) bool
	SignInUser(cnt context.Context, email, password string) (domain.User, bool)
	CreateToken(cnt context.Context, token domain.Refresh_token) bool
}
