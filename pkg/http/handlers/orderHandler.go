package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	request "github.com/Onelvay/docker-compose-project/payment/APIrequest"
	"github.com/Onelvay/docker-compose-project/payment/client"
	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/Onelvay/docker-compose-project/pkg/service"

	"github.com/google/uuid"
)

type (
	OrderHandlers struct {
		order          service.Transactioner
		db             service.ProductDbActioner
		token          service.TokenDbActioner
		userController service.UserController
	}

	OrderJSON struct {
		Product_id uint `json:"product_id"`
	}
)

func NewOrderHandler(t service.Transactioner, db service.ProductDbActioner, token service.TokenDbActioner, userController service.UserController) *OrderHandlers {
	return &OrderHandlers{t, db, token, userController}
}

func getUserIdFromBearerToken(w http.ResponseWriter, r *http.Request, s service.UserController) string {
	token, err := getTokenFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	id, err := s.ParseToken(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	return id
}
func (s *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var inp OrderJSON
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId := getUserIdFromBearerToken(w, r, s.userController)
	product, err := s.db.GetProductById(uint64(inp.Product_id))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	price := fmt.Sprintf("%v", product.Price)
	checkoutRequest := &request.CheckoutRequest{
		OrderId:           id,
		MerchantId:        "1396424",
		OrderDesc:         product.Description,
		Amount:            price,
		ProductId:         fmt.Sprint(product.Id),
		Currency:          "USD",
		ServerCallbackURL: "https://a3e1-109-239-34-71.ngrok-free.app/order/callback",
	}
	api, err := client.CreateOrder(*checkoutRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	err = s.order.CreateOrder(userId, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	js, err := json.Marshal(api)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)
}

func (s *OrderHandlers) Callback(w http.ResponseWriter, r *http.Request) { //принятие платежа
	body, _ := ioutil.ReadAll(r.Body)

	apiResp := domain.FinalResponse{}
	json.Unmarshal(body, &apiResp)
	fmt.Println(apiResp)
	s.order.CreateInfoOrder(apiResp)
}
