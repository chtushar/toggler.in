package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
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
	cookieStore *sessions.CookieStore
	jwt 				*JWT
}

//NewHandler creates a new instance of Handler
func NewHandler(log *zap.Logger, reader *request.Reader, jsonWriter *response.JSONWriter, repository *Repository, cookieStore *sessions.CookieStore, jwtSecret string) *Handler {
	return &Handler{log: log, reader: reader, jsonWriter: jsonWriter, repository: repository, cookieStore: cookieStore, jwt: NewJWT(jwtSecret)}
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
			h.log.Error("User doesn't exists", zap.Error(err))
			h.jsonWriter.NotFound(w, r, &UserNotFoundError{})
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			h.log.Error("Password doesn't match", zap.Error(err))
			h.jsonWriter.Unauthorized(w, r, &IncorrectPasswordError{})
			return
		}

		token :=	jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)

		claims["id"] = user.ID
		claims["name"] = user.Name
		claims["email"] = user.Email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, err := token.SignedString([]byte(h.jwt.Secret))

		if err != nil {
			h.log.Error("Failed to sign token", zap.Error(err))
			return
		}

		fmt.Println(tokenString)

		h.jsonWriter.Ok(w, r, &Response{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			EmailVerified: user.EmailVerified,
		})
	}
}