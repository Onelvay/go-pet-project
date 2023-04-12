package server

import (
	rest "github.com/Onelvay/docker-compose-project/pkg/controller/handleController"
	"github.com/gorilla/mux"
)

func InitRoutes(f *rest.HandleFunctions) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", f.User.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", f.User.SignIn).Methods("GET")
		auth.HandleFunc("/refresh", f.User.Refresh).Methods("GET")
	}

	books := router.PathPrefix("/books").Subrouter()
	{
		books.Use(f.User.AuthMiddleware)
		books.HandleFunc("/", f.Book.GetBooks).Methods("GET")
		books.HandleFunc("/{id}", f.Book.GetBookById).Methods("GET")
		books.HandleFunc("/{id}", f.Book.UpdateBook).Methods("PUT")
		books.HandleFunc("/{id}", f.Book.DeleteBookById).Methods("DELETE")
		router.HandleFunc("/create", f.Book.CreateBook).Methods("POST")
		router.HandleFunc("/search", f.Book.GetBooksByName).Methods("GET")

	}
	payment := router.PathPrefix("/order").Subrouter()
	{
		// payment.Use(f.AuthMiddleware)
		payment.HandleFunc("/", f.Order.CreateOrder).Methods("POST")
		payment.HandleFunc("/callback", f.Order.Callback).Methods("POST")
	}

	return router
}
