package controller

import (
	handler "github.com/Onelvay/docker-compose-project/pkg/controller/handlers"
	service "github.com/Onelvay/docker-compose-project/pkg/service"
)

type HandleFunctions struct {
	User  handler.UserHandler
	Book  handler.BookHandler
	Order handler.OrderHandlers
}

// создаем один общий хенлд класс, чтобы через него обращаться ко всем хендлерам
func NewHandlers(db service.BookstorePostgreser, userController service.UserController, or service.Transactioner, token service.TokenDbActioner) *HandleFunctions {
	b := handler.NewBookHandler(db)
	o := handler.NewOrderHandler(or, db, token)
	u := handler.NewUserHandler(userController)

	return &HandleFunctions{u, b, o}
}
