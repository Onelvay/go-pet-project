package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	service "github.com/Onelvay/docker-compose-project/pkg/service"

	"github.com/gorilla/mux"
)

type HandleFunctions struct {
	db service.BookstorePostgreser
}

func NewHandlers(db service.BookstorePostgreser) *HandleFunctions {
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
func (s *HandleFunctions) SignUp(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var inp domain.SignUpInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		panic(err)
	}
	if err := inp.Validate(); err != nil {
		panic(err)
	}
	fmt.Println(inp)
	db := s.db.(service.UserDbActioner)
	a := service.NewUserController(db)
	a.SignUp(r.Context(), inp)
	w.WriteHeader(http.StatusOK)

}

func (s *HandleFunctions) SignIn(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var inp domain.SignInInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		panic(err)
	}
	if err := inp.Validate(); err != nil {
		panic(err)
	}
	db := s.db.(service.UserDbActioner)
	a := service.NewUserController(db)
	token, err := a.SignIn(r.Context(), inp)
	if err != nil {
		panic(err)
	}
	responce, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(responce)
}

func (s *HandleFunctions) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getTokenFromRequest(r)
		if err != nil {
			panic(err)
		}
		db := s.db.(service.UserDbActioner)
		a := service.NewUserController(db)

		userId, err := a.ParseToken(r.Context(), token)
		if err != nil {
			panic(err)
		}
		ctx := context.WithValue(r.Context(), userId, userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func getTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}
	headerParts := strings.Split(header, " ")
	return headerParts[1], nil
}
