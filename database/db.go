package database

import (
	mdl "github.com/Onelvay/docker-compose-project/models"
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

func GetBooks() []mdl.Book {
	Db.Find(&books)
	return books
}

func DeleteBookById(id string) bool {
	res := Db.Where("id=?", id).Delete(&mdl.Book{})
	return res.RowsAffected == 1
}
