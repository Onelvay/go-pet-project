package controller

import (
	handler "github.com/Onelvay/docker-compose-project/pkg/handlers"
	service "github.com/Onelvay/docker-compose-project/pkg/service"
)

type HandleFunctions struct {
	Auth    handler.AuthHandler
	Product handler.ProductHandler
	Order   handler.OrderHandlers
}

// создаем один общий хенлд класс, чтобы через него обращаться ко всем хендлерам
func NewHandlers(db service.ProductDbActioner, userController *UserController, or service.Transactioner, token service.TokenDbActioner) *HandleFunctions {
	b := handler.NewProductHandler(db)
	o := handler.NewOrderHandler(or, db, token, userController)
	u := handler.NewAuthHandler(userController)

	return &HandleFunctions{u, b, o}
}
