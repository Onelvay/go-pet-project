// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type APIResponse struct {
// 	Responce interface{} `json:"responce"`
// }

// func main() {
// 	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
// 		body, _ := ioutil.ReadAll(r.Body)
// 		fmt.Println(string(body))
// 		apiResp := APIResponse{}
// 		json.Unmarshal(body, &apiResp)
// 		fmt.Println(apiResp.Responce)
// 	})
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
