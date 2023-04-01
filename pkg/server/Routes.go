package server

import (
	rest "github.com/Onelvay/docker-compose-project/pkg/rest"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	books := router.PathPrefix("/books").Subrouter()

	books.HandleFunc("/", rest.GetBooks).Methods("GET")

	books.HandleFunc("/{id}", rest.GetBookById).Methods("GET")

	books.HandleFunc("/{id}", rest.UpdateBook).Methods("PUT")

	books.HandleFunc("/{id}", rest.DeleteBookById).Methods("DELETE")

	router.HandleFunc("/create", rest.CreateBook).Methods("POST")

	router.HandleFunc("/search", rest.GetBooksByName).Methods("GET")

	return router
}
