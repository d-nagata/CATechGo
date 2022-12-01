package handler

import (
	"net/http"

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
}
