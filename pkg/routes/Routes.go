package server

import (
	rest "github.com/Onelvay/docker-compose-project/pkg/controller"
	"github.com/gorilla/mux"
)

func InitRoutes(f *rest.HandleFunctions) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/search", f.Product.GetProductsByName).Methods("GET")

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", f.Auth.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", f.Auth.SignIn).Methods("GET")
		auth.HandleFunc("/refresh", f.Auth.Refresh).Methods("GET")
	}

	user := router.PathPrefix("/profile").Subrouter()
	{
		user.HandleFunc("/orders", f.User.GetOrders).Methods("Get")
		user.HandleFunc("/orders", f.User.AddDetailToOrder).Methods("POST")
	}
	products := router.PathPrefix("/products").Subrouter()
	{
		// products.Use(f.Auth.AuthMiddleware)
		products.HandleFunc("", f.Product.GetProducts).Methods("GET")
		products.HandleFunc("/{id}", f.Product.GetProductById).Methods("GET")
		products.HandleFunc("/{id}", f.Product.DeleteProductById).Methods("DELETE")
		products.HandleFunc("/create", f.Product.CreateProduct).Methods("POST")
	}

	payment := router.PathPrefix("/order").Subrouter()
	{
		payment.HandleFunc("", f.Order.CreateOrder).Methods("POST")
		payment.HandleFunc("/callback", f.Order.Callback).Methods("POST")
	}

	return router
}
