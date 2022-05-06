package auth

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"toggler.in/internal/helpers"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
	"toggler.in/internal/proxy"
)


type Handler struct {
	log        *zap.Logger
	reader     *request.Reader
	jsonWriter *response.JSONWriter
	repository *Repository
	secureCookie *securecookie.SecureCookie
	jwt 				*helpers.JWT
}

type Config struct {
	Log 			*zap.Logger
	Reader 		*request.Reader
	JSONWriter 	*response.JSONWriter
	Repository 	*Repository
	SecureCookie *securecookie.SecureCookie
	JWTSecret 	string
}

//NewHandler creates a new instance of Handler
func NewHandler(cfg *Config) *Handler {
	return &Handler{log: cfg.Log, reader: cfg.Reader, jsonWriter: cfg.JSONWriter, repository: cfg.Repository, secureCookie: cfg.SecureCookie, jwt: helpers.NewJWT(cfg.JWTSecret)}
}


func (h *Handler) signin() http.HandlerFunc {

	type Request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8"`
	}

	type Response struct {
		ID 	 	int32 `json:"id"`
		Name 	string `json:"name"`
		Email string `json:"email"`
		EmailVerified bool `json:"emailVerified"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		ok := h.reader.ReadJSONAndValidate(w, r, req)

		if !ok {
			return
		}

		user, err := h.repository.GetUserByEmail(r.Context(),req.Email)

		if err != nil {
			h.log.Error("There was a problem logging you in", zap.Error(err))
			h.jsonWriter.NotFound(w, r, &UserNotFoundError{})
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			h.log.Error("There was a problem logging you in", zap.Error(err))
			h.jsonWriter.Unauthorized(w, r, &IncorrectPasswordError{})
			return
		}

		err = proxy.SetAuthCookie(&proxy.AuthCookieConfig{
			JWT: h.jwt,
			SC:  h.secureCookie,
			W:  &w,
			User: map[string]interface{}{
				helpers.KeyUserId: user.ID,
				helpers.KeyUserName: user.Name,
				helpers.KeyUserEmail: user.Email,
		}})

		if err != nil {
			h.log.Error("There was a problem logging you in", zap.Error(err))
				h.jsonWriter.BadRequest(w, r, &InternalError{})
			return
		}

		h.jsonWriter.Ok(w, r, &Response{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			EmailVerified: user.EmailVerified,
		})
	}
}

func (h *Handler) signout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ClearAuthCookie(&w)
		h.jsonWriter.Ok(w, r, "OK")
	}
}