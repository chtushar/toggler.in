package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

//UserRoutes adds users routes to router
func UserRoutes(r *mux.Router, uh *Handler) {
	r = r.PathPrefix("/users").Subrouter()

	r.HandleFunc("", uh.addUser()).Methods(http.MethodPost)
}