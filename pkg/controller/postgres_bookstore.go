package controller

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookstorePostgres struct {
	Db *gorm.DB
}

func NewBookstoreDbController(db *gorm.DB) *BookstorePostgres {
	return &BookstorePostgres{Db: db}
}

var book domain.Book
var books []domain.Book

func (r *BookstorePostgres) GetBookById(id string) (domain.Book, error) {
	res := r.Db.Where("id = ?", id).Find(&book)
	if res.RowsAffected == 0 {
		return domain.Book{}, errors.New("book not found")
	}
	return book, nil
}
func (r *BookstorePostgres) GetBooksByName(name string) ([]domain.Book, error) {
	res := r.Db.Where("name = ?", name).Find(&books)
	if res.RowsAffected == 0 {
		return []domain.Book{}, fmt.Errorf("no books with name %s", name)
	}
	return books, nil
}

func (r *BookstorePostgres) GetBooks(sorted bool) ([]domain.Book, error) {
	var res *gorm.DB
	if sorted {
		res = r.Db.Order("price").Find(&books)
	} else {
		res = r.Db.Order("price desc").Find(&books)
	}
	return books, res.Error
}

func (r *BookstorePostgres) DeleteBookById(id string) error {
	res := r.Db.Where("id=?", id).Delete(&domain.Book{})
	return res.Error
}
func (r *BookstorePostgres) CreateBook(name string, price float64, descr string) error {
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	res := r.Db.First(&domain.Book{}, "id = ?", id)
	if res.RowsAffected == 0 {
		res := r.Db.Create(&domain.Book{
			Id:          id,
			Name:        name,
			Description: descr,
			Price:       price,
		})
		return res.Error
	}
	return errors.New("not found")
}

func (r *BookstorePostgres) UpdateBook(id string, name string, desc string, price float64) error {
	_, res := r.GetBookById(id)
	if res == nil {
		if name != "" {
			book.Name = name
		}
		if desc != "" {
			book.Description = desc
		}
		if price != 0 {
			book.Price = price
		}
		res := r.Db.Save(&book)
		return res.Error
	}
	return res
}
