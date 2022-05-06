package handlers

import (
	"fmt"
	"net/http"
)

// https://github.com/satellity/satellity/blob/master/internal/middlewares/authentication.go

type authenticationHandler struct {}

func Authenticate(handler http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authenticating")
		handler.ServeHTTP(w, r)
	})
}