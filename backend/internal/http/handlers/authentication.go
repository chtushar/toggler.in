package handlers

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"toggler.in/internal/http/response"
)

// https://github.com/satellity/satellity/blob/master/internal/middlewares/authentication.go

type authenticationHandler struct {
	handler    http.Handler
	log        *zap.Logger
	jsonWriter *response.JSONWriter
}

func AuthenticationHandler(log *zap.Logger) func(h http.Handler) http.Handler  {
	return func(h http.Handler) http.Handler {
		ah := &authenticationHandler{
			handler:  h,
			log:      log,
		}

		return ah
	}
}

func (h *authenticationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer h.authenticate(w, r)
	h.handler.ServeHTTP(w, r)
}

func (h *authenticationHandler) authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authenticate")
}

func (h *authenticationHandler) jsonResponse(w http.ResponseWriter, r *http.Request) {
	h.jsonWriter.DefaultError(w, r)
}