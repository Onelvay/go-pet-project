package server

import (
	rest "github.com/Onelvay/docker-compose-project/pkg/controller"
	"github.com/gorilla/mux"
)

func InitRoutes(f *rest.HandleFunctions) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", f.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", f.SignIn).Methods("GET")
		auth.HandleFunc("/refresh", f.Refresh).Methods("GET")
	}

	books := router.PathPrefix("/books").Subrouter()
	{
		books.Use(f.AuthMiddleware)
		books.HandleFunc("/", f.GetBooks).Methods("GET")
		books.HandleFunc("/{id}", f.GetBookById).Methods("GET")
		books.HandleFunc("/{id}", f.UpdateBook).Methods("PUT")
		books.HandleFunc("/{id}", f.DeleteBookById).Methods("DELETE")
		router.HandleFunc("/create", f.CreateBook).Methods("POST")
		router.HandleFunc("/search", f.GetBooksByName).Methods("GET")
		router.HandleFunc("/callback", f.Callback).Methods("POST")

	}
	payment := router.PathPrefix("/order").Subrouter()
	{
		// payment.Use(f.AuthMiddleware)
		payment.HandleFunc("/", f.CreateOrder).Methods("POST")
		payment.HandleFunc("/callback", f.Callback).Methods("POST")
	}

	return router
}
