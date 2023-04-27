package server

import (
	rest "github.com/Onelvay/docker-compose-project/pkg/controller"
	"github.com/Onelvay/docker-compose-project/pkg/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes(f *rest.HandleFunctions, test handlers.UserHandler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", f.Auth.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", f.Auth.SignIn).Methods("GET")
		auth.HandleFunc("/refresh", f.Auth.Refresh).Methods("GET")
	}

	user := router.PathPrefix("/profile").Subrouter()
	{
		user.HandleFunc("", test.GetOrders).Methods("Get")
		user.HandleFunc("/orders", nil).Methods("Get")
		user.HandleFunc("/orders", nil).Methods("POST")
	}
	books := router.PathPrefix("/books").Subrouter()
	{
		// books.Use(f.Auth.AuthMiddleware)
		books.HandleFunc("", f.Book.GetBooks).Methods("GET")
		books.HandleFunc("/{id}", f.Book.GetBookById).Methods("GET")
		books.HandleFunc("/{id}", f.Book.UpdateBook).Methods("PUT")
		books.HandleFunc("/{id}", f.Book.DeleteBookById).Methods("DELETE")
		router.HandleFunc("/create", f.Book.CreateBook).Methods("POST")
		router.HandleFunc("/search", f.Book.GetBooksByName).Methods("GET")

	}
	payment := router.PathPrefix("/order").Subrouter()
	{
		//payment.Use(f.User.AuthMiddleware)
		payment.HandleFunc("", f.Order.CreateOrder).Methods("POST")
		payment.HandleFunc("/callback", f.Order.Callback).Methods("POST")
	}

	return router
}
