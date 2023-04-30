package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	userController service.UserController
}

func NewAuthHandler(userController service.UserController) *AuthHandler {
	return &AuthHandler{userController}
}
func (s *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	var inp domain.SignUpInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	if err := inp.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	err = s.userController.SignUp(r.Context(), inp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)

}
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh-token") //берем с куки рефреш токен
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	logrus.Infof("%s", cookie.Value)

	accessToken, refreshToken, err := h.userController.RefreshTokens(r.Context(), cookie.Value) //новый рефреш и bearer токен
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	responce, err := json.Marshal(map[string]string{ //форматируем в джейсон
		"token": accessToken,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	w.Header().Add("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	w.Header().Add("Content-Type", "application/json")
	w.Write(responce)

}
func (s *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	var inp domain.SignInInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	if err := inp.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	accessToken, refreshToken, err := s.userController.SignIn(r.Context(), inp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	responce, err := json.Marshal(map[string]string{
		"token": accessToken,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	w.Header().Add("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	w.Header().Add("Content-Type", "application/json")
	w.Write(responce)
}

type key int

func (s *AuthHandler) AuthMiddleware(next http.Handler) http.Handler { //проверка на то что юзер залогинился
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getTokenFromRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}
		userId, err := s.userController.ParseToken(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		var ctxUserId key
		ctx := context.WithValue(r.Context(), ctxUserId, userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func getTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		fmt.Println(headerParts)
		return "", errors.New("problems with bearer token")
	}
	return headerParts[1], nil
}
