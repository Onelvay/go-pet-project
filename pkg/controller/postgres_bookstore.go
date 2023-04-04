package controller

import (
	"context"
	"strings"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	mdl "github.com/Onelvay/docker-compose-project/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookstorePostgres struct {
	Db *gorm.DB
}

func NewDbController(db *gorm.DB) *BookstorePostgres {
	return &BookstorePostgres{Db: db}
}

var book mdl.Book
var books []mdl.Book

func (r *BookstorePostgres) GetBookById(id string) (mdl.Book, bool) {
	res := r.Db.Where("id = ?", id).Find(&book)
	if res.RowsAffected == 0 {
		return mdl.Book{}, false
	}
	return book, true
}
func (r *BookstorePostgres) GetBooksByName(name string) ([]mdl.Book, bool) {
	res := r.Db.Where("name = ?", name).Find(&books)
	if res.RowsAffected == 0 {
		return []mdl.Book{}, false
	}
	return books, true
}

func (r *BookstorePostgres) GetBooks(sorted bool) []mdl.Book {
	if sorted {
		r.Db.Order("price").Find(&books)
	} else {
		r.Db.Order("price desc").Find(&books)
	}
	return books
}

func (r *BookstorePostgres) DeleteBookById(id string) bool {
	res := r.Db.Where("id=?", id).Delete(&mdl.Book{})
	return res.RowsAffected == 1
}
func (r *BookstorePostgres) CreateBook(name string, price float64, descr string) bool {
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	res := r.Db.First(&mdl.Book{}, "id = ?", id)
	if res.RowsAffected == 0 {
		r.Db.Create(&mdl.Book{
			Id:          id,
			Name:        name,
			Description: descr,
			Price:       price,
		})
		return true
	}
	return false
}

func (r *BookstorePostgres) UpdateBook(id string, name string, desc string, price float64) bool {
	_, res := r.GetBookById(id)
	if res {
		if name != "" {
			book.Name = name
		}
		if desc != "" {
			book.Description = desc
		}
		if price != 0 {
			book.Price = price
		}
		r.Db.Save(&book)
		return true
	}
	return false
}

func (r *BookstorePostgres) Create(cnt context.Context, user domain.User) bool {
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	r.Db.Create(&domain.User{
		ID:           id,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		RegisteredAt: user.RegisteredAt,
	})
	return true
}
