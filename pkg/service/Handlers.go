package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/Onelvay/docker-compose-project/database"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	URLsort := r.URL.Query().Get("sorted")
	sort := false
	if URLsort == "true" {
		sort = true
	}
	books := db.GetBooks(sort)
	json.NewEncoder(w).Encode(books)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, _ := db.GetBookById(id)

	json.NewEncoder(w).Encode(book)
}
func GetBooksByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	books, _ := db.GetBooksByName(name)
	json.NewEncoder(w).Encode(books)
}
func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	res := db.DeleteBookById(id)
	if res {
		fmt.Fprintf(w, "успешно")
	} else {
		fmt.Fprintf(w, "не успешно")
	}
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	desc := r.URL.Query().Get("desc")
	price_str := r.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(price_str, 64)
	if name != "" && desc != "" && price != 0 && db.CreateBook(name, price, desc) {
		fmt.Fprintf(w, "успешно")
	} else {
		fmt.Fprintf(w, "не успешно")
	}
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	name := r.URL.Query().Get("name")
	desc := r.URL.Query().Get("desc")
	price_str := r.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(price_str, 64)
	res := db.UpdateBook(id, name, desc, price)
	if res {
		fmt.Fprintf(w, "успешно")
	} else {
		fmt.Fprintf(w, "не успешно")
	}

}
