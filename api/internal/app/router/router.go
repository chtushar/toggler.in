package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Routes(r *mux.Router, db *gorm.DB) {
	// HERE
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
}
