package service

import (
	mdl "github.com/Onelvay/docker-compose-project/pkg/model"
)

type Handlers interface {
	GetBooks()
	GetBookById()
	GetBookByName()
	DeleteBookById()
}
type BookstorePostgreser interface {
	GetBooks(bool) []mdl.Book
	GetBookById(string) (mdl.Book, bool)
	GetBooksByName(string) ([]mdl.Book, bool)
	DeleteBookById(string) bool
	CreateBook(string, float64, string) bool
	UpdateBook(string, string, string, float64) bool
}
