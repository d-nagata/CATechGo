package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/ikalemmon/CATechGo/handler"
	"github.com/ikalemmon/CATechGo/service"
)

func NewRouter(db *sql.DB) *mux.Router {
	// register routes
	mux := mux.NewRouter()
	mux.HandleFunc("/user", handler.NewUserHandler(service.NewUserService(db)).ServeHTTP)
	mux.HandleFunc("/gacha", handler.NewGachaHandler(service.NewGachaService(db)).ServeHTTP)
	mux.HandleFunc("/character", handler.NewCharacterHandler(service.NewCharacterService(db)).ServeHTTP)
	return mux
}
