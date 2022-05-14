package handlers

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type corsHandler struct {
	handler    http.Handler
	log        *zap.Logger
}

func CORSHandler(log *zap.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		ch := &corsHandler{
			handler: h,
			log: log,
		}

		return ch
	}
}

func (h *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	defer h.checkCorsHeaders(w, r)
}

func (h *corsHandler) checkCorsHeaders(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path, r.Method)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,DELETE")
		w.Header().Set("Access-Control-Max-Age", "86400")
		h.handler.ServeHTTP(w, r)
}