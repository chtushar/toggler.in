package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AuthRoutes adds auth routes to router
func AuthRoutes(r *mux.Router, ah *Handler) {
	r = r.PathPrefix("/auth").Subrouter()

	r.HandleFunc("/signin", ah.signin()).Methods(http.MethodGet)
}
