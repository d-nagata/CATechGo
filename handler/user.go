package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ikalemmon/CATechGo/model"
	"github.com/ikalemmon/CATechGo/service"
)

// A UserHandler implements handling REST endpoints.
type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { //ユーザー情報作成
		request := &model.UserCreateRequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Println(err)
			return
		}
		if request.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token, err := h.svc.CreateUser(r.Context(), request.Name)
		if err != nil {
			log.Println(err)
			return
		}
		response := model.UserCreateResponse{Token: token}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err)
			return
		}
	}
	if r.Method == "GET" { //ユーザー情報取得
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		if tokenString == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		name, err := h.svc.GetUser(r.Context(), tokenString)
		if err != nil {
			log.Println(err)
			return
		}
		response := model.UserGetResponse{Name: name}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err)
			return
		}
	}
	if r.Method == "PUT" { //ユーザー情報更新
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		request := &model.UserUpdateRequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Println(err)
			return
		}
		if request.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := h.svc.UpdateUser(r.Context(), tokenString, request.Name)
		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

}
