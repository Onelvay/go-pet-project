package controller

import (
	"github.com/Onelvay/docker-compose-project/pkg/http/handlers"
	service "github.com/Onelvay/docker-compose-project/pkg/service"
)

type HandleFunctions struct {
	Auth    *handlers.AuthHandler
	Product *handlers.ProductHandler
	Order   *handlers.OrderHandlers
	User    *handlers.UserHandler
}

func NewHandlers(productDb service.ProductDbActioner, userController *UserController, or service.Transactioner, token service.TokenDbActioner, userDbController service.UserDbActioner) *HandleFunctions {
	product := handlers.NewProductHandler(productDb)
	order := handlers.NewOrderHandler(or, productDb, token, userController)
	auth := handlers.NewAuthHandler(userController)
	user := handlers.NewUserHandler(userController, userDbController, productDb)

	return &HandleFunctions{auth, product, order, user}
}
