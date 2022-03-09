package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Routes(r *mux.Router, db *gorm.DB, log *zap.Logger,) {
	// HERE
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
}
