package controller

import (
	request "github.com/Onelvay/docker-compose-project/payment/APIrequest"
	handler "github.com/Onelvay/docker-compose-project/pkg/controller/handlers"
	service "github.com/Onelvay/docker-compose-project/pkg/service"
)

type Transactioner interface {
	CreateOrder(userId string, orderId string)
	CreateInfoOrder(request.FinalResponse)
}

type HandleFunctions struct {
	User  handler.UserHandler
	Book  handler.BookHandler
	Order handler.OrderHandlers
}

func NewHandlers(db service.BookstorePostgreser, userController service.UserController, or Transactioner) *HandleFunctions {
	b := handler.NewBookHandler(db)
	o := handler.NewOrderHandler(or, db)
	u := handler.NewUserHandler(userController)

	return &HandleFunctions{u, b, o}
}
