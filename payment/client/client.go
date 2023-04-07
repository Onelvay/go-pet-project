package client

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	req "github.com/Onelvay/docker-compose-project/payment/request"
	"github.com/fatih/structs"
	"github.com/google/uuid"
)

const (
	checkoutUrl      = "https://pay.fondy.eu/api/checkout/url/"
	merchantId       = "1396424"
	merchantPassword = "test"
	currency         = "USD"
	language         = "ru"
)

type CheckoutRequest struct {
	OrderId           string `json:"order_id"`
	MerchantId        string `json:"merchant_id"`
	OrderDesc         string `json:"order_desc"`
	Signature         string `json:"signature"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	ResponseURL       string `json:"response_url,omitempty"`
	ServerCallbackURL string `json:"server_callback_url,omitempty"`
	SenderEmail       string `json:"sender_email,omitempty"`
	Language          string `json:"lang,omitempty"`
	ProductId         string `json:"product_id,omitempty"`
}

func (r *CheckoutRequest) SetSignature(password string) {
	params := structs.Map(r)
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	values := []string{}
	for _, key := range keys {
		value := params[key].(string)
		if value == "" {
			continue
		}
		values = append(values, value)
	}
	r.Signature = generateSignature(values, password)
}
func generateSignature(values []string, password string) string {
	newVals := []string{password}
	newVals = append(newVals, values...)
	signatureString := strings.Join(newVals, "|")
	fmt.Println(signatureString)
	hash := sha1.New()
	hash.Write([]byte(signatureString))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func CreateOrder(amount string) {
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	checkoutRequest := &CheckoutRequest{
		OrderId:           id,
		MerchantId:        merchantId,
		OrderDesc:         "course fsafx aaa",
		Amount:            amount,
		Currency:          "USD",
		ServerCallbackURL: "https://6a8f-80-242-211-178.in.ngrok.io/callback",
	}
	checkoutRequest.SetSignature(merchantPassword)
	request := req.APIRequest{Request: checkoutRequest}
	requestBody, _ := json.Marshal(request)
	resp, err := http.Post(checkoutUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	apiResp := req.APIResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(apiResp)
}
