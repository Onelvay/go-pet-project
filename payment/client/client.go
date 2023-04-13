package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	req "github.com/Onelvay/docker-compose-project/payment/APIrequest"
)

var merchantId, merchantPassword, checkoutUrl string

func InitConst(merchantId_, merchantPassword_, checkoutUrl_ string) {
	merchantId = merchantId_
	merchantPassword = merchantPassword_
	checkoutUrl = checkoutUrl_
}

func CreateOrder(checkoutRequest req.CheckoutRequest) (req.APIResponse, error) {
	checkoutRequest.SetSignature(merchantPassword)
	request := req.APIRequest{Request: checkoutRequest}
	requestBody, _ := json.Marshal(request)
	resp, err := http.Post(checkoutUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return req.APIResponse{}, err
	}
	defer resp.Body.Close()

	apiResp := req.APIResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return req.APIResponse{}, err
	}

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return req.APIResponse{}, err
	}
	return apiResp, nil
}
