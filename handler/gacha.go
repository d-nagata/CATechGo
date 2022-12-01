package handler

import (
	"net/http"

	"github.com/ikalemmon/CATechGo/service"
)

// A GachaHandler implements handling REST endpoints.
type GachaHandler struct {
	svc *service.GachaService
}

func NewGachaHandler(svc *service.GachaService) *GachaHandler {
	return &GachaHandler{
		svc: svc,
	}
}

func (h *GachaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
