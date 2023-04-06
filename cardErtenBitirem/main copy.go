// package main

// import (
// 	"bytes"
// 	"crypto/sha1"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"sort"
// 	"strings"

// 	"github.com/fatih/structs"
// )

// const (
// 	checkoutUrl      = "https://pay.fondy.eu/api/checkout/url/"
// 	merchantId       = "1396424"
// 	merchantPassword = "test"
// 	currency         = "USD"
// 	language         = "ru"
// )

// type APIRequest struct {
// 	Request interface{} `json:"request"`
// }
// type APIResponse struct {
// 	Response interface{} ` json:"response"`
// }
// type CheckoutRequest struct {
// 	OrderId           string `json:"order_id"`
// 	MerchantId        string `json:"merchant_id"`
// 	OrderDesc         string `json:"order_desc"`
// 	Signature         string `json:"signature"`
// 	Amount            string `json:"amount"`
// 	Currency          string `json:"currency"`
// 	ResponseURL       string `json:"response_url, omitempty"`
// 	ServerCallbackURL string `json:"server_callback_url, omitempty"`
// 	SenderEmail       string `json:"sender_email, omitempty"`
// 	Language          string `json:"lang, omitempty"`
// 	ProductId         string `json: "product_id, omitempty"`
// }

// func (r *CheckoutRequest) SetSignature(password string) {
// 	params := structs.Map(r)
// 	var keys []string
// 	for k := range params {
// 		keys = append(keys, k)
// 	}
// 	sort.Strings(keys)
// 	values := []string{}
// 	for _, key := range keys {
// 		value := params[key].(string)
// 		if value == "" {
// 			continue
// 		}
// 		values = append(values, value)
// 	}
// 	r.Signature = generateSignature(values, password)
// }
// func generateSignature(values []string, password string) string {
// 	newVals := []string{password}
// 	newVals = append(newVals, values...)
// 	signatureString := strings.Join(newVals, "|")
// 	fmt.Println(signatureString)
// 	hash := sha1.New()
// 	hash.Write([]byte(signatureString))
// 	return fmt.Sprintf("%x", hash.Sum(nil))
// }
// func main() {
// 	id := "asdasdasszxefwdtt466322322w"
// 	req := &CheckoutRequest{
// 		OrderId:           id,
// 		MerchantId:        merchantId,
// 		OrderDesc:         "course fsafx aaa",
// 		Amount:            "1200",
// 		Currency:          "USD",
// 		ServerCallbackURL: "https://6a8f-80-242-211-178.in.ngrok.io/callback",
// 	}
// 	req.SetSignature(merchantPassword)
// 	request := APIRequest{Request: req}
// 	requestBody, _ := json.Marshal(request)
// 	resp, err := http.Post(checkoutUrl, "application/json", bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	apiResp := APIResponse{}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = json.Unmarshal(body, &apiResp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("avavasdas", apiResp, "zxczxcxzczx")
// }
