package handlers

import (
	"fmt"
	"net/http"

	"github.com/Onelvay/docker-compose-project/pkg/service"
)

type UserHandler struct {
	userController service.UserController
	userDb         service.UserDbActioner
}

func NewUserHandler(userController service.UserController, userDb service.UserDbActioner) *UserHandler {
	return &UserHandler{userController, userDb}
}

func (u *UserHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userId := getUserIdFromBearerToken(w, r, u.userController)
	res, err := u.userDb.GetUserOrders(userId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
