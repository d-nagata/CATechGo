package handler

import (
	"net/http"

	"github.com/ikalemmon/CATechGo/service"
)

// A CharacterHandler implements handling REST endpoints.
type CharacterHandler struct {
	svc *service.CharacterService
}

func NewCharacterHandler(svc *service.CharacterService) *CharacterHandler {
	return &CharacterHandler{
		svc: svc,
	}
}

func (h *CharacterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
