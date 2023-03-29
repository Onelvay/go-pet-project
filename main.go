package main

import (
	"time"

	// "net/http"
	// "os"

	// "github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	Id          string    `gorm:"primary_key" json:"id"`
	Created_at  time.Time `gorm:"default:current_timestamp"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"desc"`
}

// func homePage(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("asdasd")
// 	fmt.Fprintf(w, "this is homepage")
// }

func main() {
	dsn := "host=localhost user=postgres password=Adg12332, dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// var book Book
	// var books []Book
	// Db.Raw("SELECT * FROM books LIMIT 1").Scan(&book) //vse knigi limit ybrat'
	// Db.Create(&Book{Id: "D2212", Price: 100, Created_at: time.Now(), Name: "alenov", Description: "asdas adm qowe kakfk zkkx kckdk kkkkk"})
	// Db.Find(&books)
	// fmt.Println(book) // Read
	// Db.Where("id = ?", "D2222").First(&book)
	// fmt.Println(books)
	// Db.Where("id=?", "D2212").Delete(&book)
	// var PORT string
	// if PORT = os.Getenv("PORT"); PORT == "" {
	// 	PORT = "8080"
	// }
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homePage)
	// err := http.ListenAndServe(":"+PORT, router)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

}
