package server

import (
	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/gorilla/mux"
)

type Controller struct {
	services *service.Service
}

func NewController(services *service.Service) *Controller {
	return &Controller{services: services}
}

func InitRoutes(c *Controller) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", c.services.HomePage)
	// books := router.PathPrefix("/books").Subrouter()
	// books.HandleFunc("/{id}", c.services.HomePage)
	return router
}
