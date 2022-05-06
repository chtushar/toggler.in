package auth

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
)


type Handler struct {
	log        *zap.Logger
	reader     *request.Reader
	jsonWriter *response.JSONWriter
	repository *Repository
	secureCookie *securecookie.SecureCookie
	jwt 				*JWT
}

//NewHandler creates a new instance of Handler
func NewHandler(log *zap.Logger, reader *request.Reader, jsonWriter *response.JSONWriter, repository *Repository, secureCookie *securecookie.SecureCookie, jwtSecret string) *Handler {
	return &Handler{log: log, reader: reader, jsonWriter: jsonWriter, repository: repository, secureCookie: secureCookie, jwt: NewJWT(jwtSecret)}
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

		accessTokenData := map[string]interface{}{
			KeyUserId: user.ID,
			KeyUserEmail: user.Email,
			KeyUserName: user.Name,
		}

		token, err :=	h.jwt.NewToken(AuthSecret, accessTokenData)

		if err != nil {
			h.log.Error("Failed to sign token", zap.Error(err))
				h.jsonWriter.Unauthorized(w, r, &IncorrectPasswordError{})
			return
		}

		cookieCoded, err := h.secureCookie.Encode("auth", token)

		if err != nil {
			h.log.Error("Cookie coding error", zap.Error(err))
			h.jsonWriter.Unauthorized(w, r, &IncorrectPasswordError{})
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth",
			Value:    cookieCoded,
			Path:     "/",
			HttpOnly: true,
		})

		h.jsonWriter.Ok(w, r, &Response{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			EmailVerified: user.EmailVerified,
		})
	}
}