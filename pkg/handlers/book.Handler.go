package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/gorilla/mux"
)

type BookHandler struct {
	db service.BookstorePostgreser
}

func NewBookHandler(db service.BookstorePostgreser) BookHandler {
	return BookHandler{db}
}

func (s *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) { //ниже все понятно думаю
	URLsort := r.URL.Query().Get("sorted")
	sort := false
	if URLsort == "true" {
		sort = true
	}
	books, err := s.db.GetBooks(sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(books)
}
func (s *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, err := s.db.GetBookById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	json.NewEncoder(w).Encode(book)
}
func (s *BookHandler) GetBooksByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	books, err := s.db.GetBooksByName(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(books)
}
func (s *BookHandler) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	res := s.db.DeleteBookById(id)
	if res == nil {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
func (s *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	desc := r.URL.Query().Get("desc")
	price_str := r.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(price_str, 64)
	if name != "" && desc != "" && price != 0 && s.db.CreateBook(name, price, desc) == nil {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	name := r.URL.Query().Get("name")
	desc := r.URL.Query().Get("desc")
	price_str := r.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(price_str, 64)
	res := s.db.UpdateBook(id, name, desc, price)
	if res == nil {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
