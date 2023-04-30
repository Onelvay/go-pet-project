package controller

import (
	handler "github.com/Onelvay/docker-compose-project/pkg/handlers"
	service "github.com/Onelvay/docker-compose-project/pkg/service"
)

type HandleFunctions struct {
	Auth    *handler.AuthHandler
	Product *handler.ProductHandler
	Order   *handler.OrderHandlers
	User    *handler.UserHandler
}

func NewHandlers(productDb service.ProductDbActioner, userController *UserController, or service.Transactioner, token service.TokenDbActioner, userDbController service.UserDbActioner) *HandleFunctions {
	product := handler.NewProductHandler(productDb)
	order := handler.NewOrderHandler(or, productDb, token, userController)
	auth := handler.NewAuthHandler(userController)
	user := handler.NewUserHandler(userController, userDbController, productDb)

	return &HandleFunctions{auth, product, order, user}
}
