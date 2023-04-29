package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	db service.ProductDbActioner
}

func NewProductHandler(db service.ProductDbActioner) ProductHandler {
	return ProductHandler{db}
}

func (s *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) { //ниже все понятно думаю
	URLsort := r.URL.Query().Get("sorted")
	sort := false
	if URLsort == "true" {
		sort = true
	}
	products, err := s.db.GetProducts(sort)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(products)
}
func (s *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uintid, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	}
	product, err := s.db.GetProductById(uintid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}
func (s *ProductHandler) GetProductsByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	products, err := s.db.GetProductsByName(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(products)
}
func (s *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	var inp domain.Product
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	s.db.CreateProduct(inp)
}

// func (s *ProductHandler) DeleteBookById(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]
// 	res := s.db.DeleteBookById(id)
// 	if res == nil {
// 		w.WriteHeader(http.StatusAccepted)
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// }

// func (s *ProductHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]
// 	name := r.URL.Query().Get("name")
// 	desc := r.URL.Query().Get("desc")
// 	price_str := r.URL.Query().Get("price")
// 	price, _ := strconv.ParseFloat(price_str, 64)
// 	res := s.db.UpdateBook(id, name, desc, price)
// 	if res == nil {
// 		w.WriteHeader(http.StatusAccepted)
// 	} else {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// }
