package server

import (
	rest "github.com/Onelvay/docker-compose-project/pkg/controller"
	"github.com/gorilla/mux"
)

func InitRoutes(f *rest.HandleFunctions) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	books := router.PathPrefix("/books").Subrouter()

	books.HandleFunc("/", f.GetBooks).Methods("GET")

	books.HandleFunc("/{id}", f.GetBookById).Methods("GET")

	books.HandleFunc("/{id}", f.UpdateBook).Methods("PUT")

	books.HandleFunc("/{id}", f.DeleteBookById).Methods("DELETE")

	router.HandleFunc("/create", f.CreateBook).Methods("POST")

	router.HandleFunc("/search", f.GetBooksByName).Methods("GET")

	return router
}
