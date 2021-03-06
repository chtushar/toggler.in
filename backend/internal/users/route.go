package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

//UserRoutes adds users routes to router
func UserRoutes(r *mux.Router, uh *Handler) {
	r = r.PathPrefix("/users").Subrouter()

	r.HandleFunc("/signup", uh.addUser()).Methods(http.MethodPost)
	r.HandleFunc("/user_info", uh.getUser()).Methods(http.MethodGet)
}