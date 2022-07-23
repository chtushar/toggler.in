package handlers

import (
	"net/http"

	"github.com/rs/cors"
	"go.uber.org/zap"
)

type corsHandler struct {
	handler    http.Handler
	log        *zap.Logger
}

func CORSHandler(log *zap.Logger) func(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug: true,
	})

	return func(h http.Handler) http.Handler {
		ch := &corsHandler{
			handler: c.Handler(h),
			log: log,
		}

		return ch
	}
}

func (h *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	h.handler.ServeHTTP(w, r)
}
