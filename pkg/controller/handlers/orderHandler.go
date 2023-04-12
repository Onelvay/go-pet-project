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

type Transactioner interface {
	CreateOrder(userId string, orderId string)
	CreateInfoOrder(request.FinalResponse)
}

type OrderHandlers struct {
	order Transactioner
	db    service.BookstorePostgreser
}

func NewOrderHandler(t Transactioner, db service.BookstorePostgreser) OrderHandlers {
	return OrderHandlers{t, db}
}

type Transaction struct {
	Product_id string
	User_id    string //временно
}

func (s *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var inp Transaction
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		panic(err)
	}
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
	s.order.CreateOrder(inp.User_id, id)
	json.NewEncoder(w).Encode(api)

}

func (s *OrderHandlers) Callback(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	apiResp := request.FinalResponse{}
	json.Unmarshal(body, &apiResp)
	s.order.CreateInfoOrder(apiResp)
}
