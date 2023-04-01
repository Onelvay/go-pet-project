package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	contr "github.com/Onelvay/docker-compose-project/pkg/controller"

	"github.com/gorilla/mux"
)

type HandleFunctions struct {
	db *contr.BookstorePostgres
}

func NewHandlers(db *contr.BookstorePostgres) *HandleFunctions {
	return &HandleFunctions{db}
}

func (s *HandleFunctions) GetBooks(w http.ResponseWriter, r *http.Request) {
	URLsort := r.URL.Query().Get("sorted")
	sort := false
	if URLsort == "true" {
		sort = true
	}
	books := s.db.GetBooks(sort)
	json.NewEncoder(w).Encode(books)
}
func (s *HandleFunctions) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, _ := s.db.GetBookById(id)

	json.NewEncoder(w).Encode(book)
}
func (s *HandleFunctions) GetBooksByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	books, _ := s.db.GetBooksByName(name)
	json.NewEncoder(w).Encode(books)
}
func (s *HandleFunctions) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	res := s.db.DeleteBookById(id)
	if res {
		fmt.Fprintf(w, "успешно")
	} else {
		fmt.Fprintf(w, "не успешно")
	}
}
func (s *HandleFunctions) CreateBook(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	desc := r.URL.Query().Get("desc")
	price_str := r.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(price_str, 64)
	if name != "" && desc != "" && price != 0 && s.db.CreateBook(name, price, desc) {
		fmt.Fprintf(w, "успешно")
	} else {
		fmt.Fprintf(w, "не успешно")
	}
}
func (s *HandleFunctions) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	name := r.URL.Query().Get("name")
	desc := r.URL.Query().Get("desc")
	price_str := r.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(price_str, 64)
	res := s.db.UpdateBook(id, name, desc, price)
	if res {
		fmt.Fprintf(w, "успешно")
	} else {
		fmt.Fprintf(w, "не успешно")
	}

}
