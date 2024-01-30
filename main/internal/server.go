package internal

import (
	"net/http"

	"github.com/artHouse-Docs/backend/internal/user"
	"github.com/gorilla/mux"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	r.Path("/users").HandlerFunc(user.RegisterUser).Methods(http.MethodPost)
	r.Path("/users/login").HandlerFunc(user.LoginUser).Methods(http.MethodPost)

	return r
}
