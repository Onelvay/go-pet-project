package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Println(err)
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
func (u *UserHandler) AddDetailToOrder(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	userId := getUserIdFromBearerToken(w, r, u.userController)
	var req domain.OrderDetail
	if err = json.Unmarshal(bytes, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	req.User_id = userId
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	err = u.userDb.AddDetailToOrder(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
