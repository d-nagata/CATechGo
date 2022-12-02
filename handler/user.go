package handler

import (
	"encoding/json"
	"log"
	"net/http"

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
	}
	if r.Method == "PUT" { //ユーザー情報更新
	}
}
