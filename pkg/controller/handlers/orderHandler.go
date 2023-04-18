package handlers

//все что ниже трогать не стоит

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	request "github.com/Onelvay/docker-compose-project/payment/APIrequest"
	"github.com/Onelvay/docker-compose-project/payment/client"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/google/uuid"
)

var Mutex sync.Mutex

type OrderHandlers struct {
	order          service.Transactioner
	db             service.BookstorePostgreser
	token          service.TokenDbActioner
	userController service.UserController
}

func NewOrderHandler(t service.Transactioner, db service.BookstorePostgreser, token service.TokenDbActioner, userController service.UserController) OrderHandlers {
	return OrderHandlers{t, db, token, userController}
}

type OrderJSON struct {
	Product_id string
}

func (s *OrderHandlers) findUser(w http.ResponseWriter, r *http.Request, userId chan string) {
	Mutex.Lock()
	token, err := getTokenFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("asdas")
		panic(err)
	}
	Id, err := s.userController.ParseToken(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	userId <- Id
	Mutex.Unlock()
}
func (s *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "123121312")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var inp OrderJSON
	if err = json.Unmarshal(reqBytes, &inp); err != nil { //принятие данных джейсон
		fmt.Println(err, "4534534345")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := make(chan string)
	go s.findUser(w, r, userId)

	product, err := s.db.GetBookById(inp.Product_id)
	if err != nil {
		fmt.Println(err)
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
		ProductId:         product.Id,
		Currency:          "USD",
		ServerCallbackURL: "https://0939-80-242-211-178.in.ngrok.io/order/callback",
	}
	api, err := client.CreateOrder(*checkoutRequest) //отправляем запрос на заказ
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	err = s.order.CreateOrder(<-userId, id) //создаем в дб заказ
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
	fmt.Println(apiResp)
	s.order.CreateInfoOrder(apiResp)
}
