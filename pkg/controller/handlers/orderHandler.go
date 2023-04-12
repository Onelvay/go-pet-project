package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	request "github.com/Onelvay/docker-compose-project/payment/APIrequest"
	"github.com/Onelvay/docker-compose-project/payment/client"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/google/uuid"
)

type OrderHandlers struct {
	order service.Transactioner
	db    service.BookstorePostgreser
	token service.TokenPostgreser
}

func NewOrderHandler(t service.Transactioner, db service.BookstorePostgreser, token service.TokenPostgreser) OrderHandlers {
	return OrderHandlers{t, db, token}
}

type OrderJSON struct {
	Product_id string
}

func (s *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var inp OrderJSON
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		panic(err)
	}
	cookie, err := r.Cookie("refresh-token")
	if err != nil {
		panic(err)
	}
	userId := s.token.GetUserIdByToken(cookie.Value)
	product, _ := s.db.GetBookById(inp.Product_id)
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	price := fmt.Sprintf("%v", product.Price)
	checkoutRequest := &request.CheckoutRequest{
		OrderId:           id,
		MerchantId:        "1396424",
		OrderDesc:         product.Description,
		Amount:            price,
		ProductId:         product.Id,
		Currency:          "USD",
		ServerCallbackURL: "https://8d8d-80-242-211-178.in.ngrok.io/callback",
	}
	api := client.CreateOrder(*checkoutRequest)
	s.order.CreateOrder(userId, id)
	json.NewEncoder(w).Encode(api)

}

func (s *OrderHandlers) Callback(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	apiResp := request.FinalResponse{}
	json.Unmarshal(body, &apiResp)
	s.order.CreateInfoOrder(apiResp)
}
