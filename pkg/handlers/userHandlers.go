package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/Onelvay/docker-compose-project/pkg/service"
)

type UserHandler struct {
	userController service.UserController
	userDb         service.UserDbActioner
	productDb      service.ProductDbActioner
}

func NewUserHandler(userController service.UserController, userDb service.UserDbActioner, productDb service.ProductDbActioner) *UserHandler {
	return &UserHandler{userController, userDb, productDb}
}

func (u *UserHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userId := getUserIdFromBearerToken(w, r, u.userController)
	productIds, err := u.userDb.GetUserOrders(userId)
	if err != nil {
		fmt.Println(err)
	}
	var products []domain.Product
	for i := 0; i < len(productIds); i++ {
		product, err := u.productDb.GetProductById(uint64(productIds[i]))
		if err != nil {
			continue
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}
