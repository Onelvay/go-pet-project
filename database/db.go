package database

import (
	"strings"

	mdl "github.com/Onelvay/docker-compose-project/pkg/model"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error
var book mdl.Book
var books []mdl.Book

func init() {
	dsn := "host=localhost user=postgres password=Adg12332, dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetBookById(id string) (mdl.Book, bool) {
	res := Db.First(&book, "id = ?", id)
	if res.RowsAffected == 0 {
		return mdl.Book{}, false
	}
	return book, true
}
func GetBooksByName(name string) ([]mdl.Book, bool) {
	res := Db.First(&books, "name = ?", name)
	if res.RowsAffected == 0 {
		return []mdl.Book{}, false
	}
	return books, true
}

func GetBooks(sorted bool) []mdl.Book {
	if sorted {
		Db.Order("price").Find(&books)
	} else {
		Db.Order("price desc").Find(&books)
	}
	return books
}

func DeleteBookById(id string) bool {
	res := Db.Where("id=?", id).Delete(&mdl.Book{})
	return res.RowsAffected == 1
}
func CreateBook(name string, price float64, descr string) bool {
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	res := Db.First(&mdl.Book{}, "id = ?", id)
	if res.RowsAffected == 0 {
		Db.Create(&mdl.Book{
			Id:          id,
			Name:        name,
			Description: descr,
			Price:       price,
		})
		return true
	}
	return false
}

func UpdateNameOfBook(id string, name string) bool {
	res := Db.First(book, "id = ?", id)
	if res.RowsAffected == 1 {
		book.Name = name
		Db.Save(&book)
		return true
	}
	return false
}
func UpdateDescrOfBook(id string, descr string) bool {
	res := Db.First(book, "id = ?", id)
	if res.RowsAffected == 1 {
		book.Description = descr
		Db.Save(&book)
		return true
	}
	return false
}
