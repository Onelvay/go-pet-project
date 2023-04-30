package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	db service.ProductDbActioner
}

func NewProductHandler(db service.ProductDbActioner) *ProductHandler {
	return &ProductHandler{db}
}

func (s *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) { //ниже все понятно думаю
	URLsort := r.URL.Query().Get("sorted")
	sort := false
	if URLsort == "true" {
		sort = true
	}
	products, err := s.db.GetProducts(sort)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
func (s *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uintid, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}
	product, err := s.db.GetProductById(uintid)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}
func (s *ProductHandler) GetProductsByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	products, err := s.db.GetProductsByName(name)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(products)
}
func (s *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	now := time.Now()
	seconds := now.Unix()
	uintTime := uint64(seconds)
	if err != nil {
		panic(err)
	}
	var inp domain.Product
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	inp.Id = uintTime
	s.db.CreateProduct(inp)
}

func (s *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uintId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}
	err = s.db.DeleteProductById(uintId)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
