package service

import (
	"context"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
)

type PasswordHasher interface {
	Hash(password string) string
}
type UserDbActioner interface {
	CreateUser(cnt context.Context, user domain.User) error
	SignInUser(context.Context, string, string) (domain.User, error)
	GetUserOrders(id string) ([]domain.UserOrders, error)
}
type TokenDbActioner interface {
	CreateToken(cnt context.Context, token domain.Refresh_token) error
	GetToken(cxt context.Context, token string) (domain.Refresh_token, error)
	GetUserIdByToken(token string) (string, error)
}

type Transactioner interface {
	CreateOrder(userId string, orderId string) error
	CreateInfoOrder(domain.FinalResponse) error
}

type ProductPostgreser interface {
	GetBooks(bool) ([]domain.Product, error)
	GetBookById(string) (domain.Product, error)
	GetBooksByName(string) ([]domain.Product, error)
	DeleteBookById(string) error
	CreateBook(string, float64, string) error
	UpdateBook(string, string, string, float64) error
}
type UserController interface {
	SignUp(ctx context.Context, inp domain.SignUpInput) error
	SignIn(ctx context.Context, inp domain.SignInInput) (string, string, error)
	ParseToken(ctx context.Context, token string) (string, error)
	RefreshTokens(ctx context.Context, refreshToken string) (string, string, error)
	GenerateTokens(ctx context.Context, userId string) (string, string, error)
}

// type Seller interface {
// 	CreateBook(domain.Book) error
// }
