package handlers

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"go.uber.org/zap"
	"toggler.in/internal/helpers"
	"toggler.in/internal/http/response"
)

// https://github.com/satellity/satellity/blob/master/internal/middlewares/authentication.go

type authenticationHandler struct {
	handler    http.Handler
	log        *zap.Logger
	jsonWriter *response.JSONWriter
	sc 	*securecookie.SecureCookie
	jwt *helpers.JWT
}

var whiteList = [][2]string {
	{"/api/v1/auth/signin", "POST"},
	{"/api/v1/auth/signout", "POST"},
	{"/api/v1/users/signup", "POST"},
}

func AuthenticationHandler(log *zap.Logger, sc *securecookie.SecureCookie, jwt *helpers.JWT) func(h http.Handler) http.Handler  {
	return func(h http.Handler) http.Handler {
		ah := &authenticationHandler{
			handler:  h,
			log:  		log,
			sc: 			sc,
			jwt: 			jwt,
		}

		return ah
	}
}

func (h *authenticationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer h.authenticate(w, r)
}

func (h *authenticationHandler) authenticate(w http.ResponseWriter, r *http.Request) {

	for _, white := range whiteList {
		if r.URL.Path == white[0] && r.Method == white[1] {
			h.handler.ServeHTTP(w, r)
			return
		}
	}

	cookie, err := r.Cookie("auth");


	if err != nil {
		h.jsonWriter.Unauthorized(w, r, &UnauthorizedError{})
		h.log.Error("There was a problem authenticating", zap.Error(err))
		return
	}

	var decodedCookieValue string;
	err = h.sc.Decode("auth", cookie.Value, &decodedCookieValue)

	if err != nil {
		h.jsonWriter.Unauthorized(w, r, &UnauthorizedError{})
		h.log.Error("There was a problem authenticating", zap.Error(err))
		return
	}


	_, err = h.jwt.ReadTokenAndValidate(decodedCookieValue)

	if err != nil  {
			h.jsonWriter.Unauthorized(w, r, &UnauthorizedError{})
			h.log.Error("There was a problem authenticating", zap.Error(err))
			return
	}


	h.handler.ServeHTTP(w, r)
}
