package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

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
	querySort := r.URL.Query().Get("sort")
	products, err := s.db.GetProducts()
	for i, v := range products {
		rating := s.db.GetProductRating(uint(v.Id))
		products[i].Rating = rating
	}
	if querySort == "rating" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Rating > products[j].Rating
		})
	} else if querySort == "price" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price > products[j].Price
		})
	}

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)
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
	js, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)

}
func (s *ProductHandler) GetProductsByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	products, err := s.db.GetProductsByName(name)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusNotFound)
		return
	}
	js, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)
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
