package main

import (

	// "net/http"
	// "os"

	// "github.com/gorilla/mux"

	"fmt"

	db "github.com/Onelvay/docker-compose-project/database"
)

// func homePage(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("asdasd")
// 	fmt.Fprintf(w, "this is homepage")
// }

func main() {
	// res, ans := db.GetBookById("d2222")
	// fmt.Println(res, ans)

	// res := db.GetBooks()
	// fmt.Println(res)
	// db.DeleteBookById("11")
	fmt.Println(db.GetBooksByName("asd"))
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
