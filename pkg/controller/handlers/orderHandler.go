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
	token service.TokenDbActioner
}

func NewOrderHandler(t service.Transactioner, db service.BookstorePostgreser, token service.TokenDbActioner) OrderHandlers {
	return OrderHandlers{t, db, token}
}

type OrderJSON struct {
	Product_id string
}

func (s *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var inp OrderJSON
	if err = json.Unmarshal(reqBytes, &inp); err != nil { //принятие данных джейсон
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("refresh-token") //читаем с куки токен
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := s.token.GetUserIdByToken(cookie.Value) //находим юзера
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	api, err := client.CreateOrder(*checkoutRequest) //отправляем запрос на заказ
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	err = s.order.CreateOrder(userId, id) //создаем в дб заказ
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	json.NewEncoder(w).Encode(api) //ретерним ссылку на оплату

}

func (s *OrderHandlers) Callback(w http.ResponseWriter, r *http.Request) { //принятие платежа
	body, _ := ioutil.ReadAll(r.Body) //читаем джейсон

	apiResp := request.FinalResponse{}
	json.Unmarshal(body, &apiResp)
	s.order.CreateInfoOrder(apiResp)
}
