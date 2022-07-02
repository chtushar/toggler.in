package teams

import (
	"net/http"

	"github.com/gorilla/mux"
)

func TeamRoutes(r *mux.Router, uh *Handler)  {
	r = r.PathPrefix("/teams").Subrouter()

	r.HandleFunc("/create", uh.createTeam()).Methods(http.MethodPost)
}