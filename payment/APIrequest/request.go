package request

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/structs"
)

type APIRequest struct {
	Request interface{} `json:"request"`
}
type APIResponse struct {
	Response interface{} `json:"response"`
}
type APIResponce struct {
	Responce interface{} `json:"responce"`
}
type CheckoutRequest struct {
	Sender_account    string `json:"sender_account"`
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

	hash := sha1.New()
	hash.Write([]byte(signatureString))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
