package server

import (
	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	books := router.PathPrefix("/books").Subrouter()

	books.HandleFunc("/", service.GetBooks).Methods("GET")

	books.HandleFunc("/{id}", service.GetBookById).Methods("GET")

	books.HandleFunc("/{id}", service.UpdateBook).Methods("PUT")

	books.HandleFunc("/{id}", service.DeleteBookById).Methods("DELETE")

	router.HandleFunc("/create", service.CreateBook).Methods("POST")

	router.HandleFunc("/search", service.GetBooksByName).Methods("GET")

	return router
}
