package users

import (
	"net/http"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"toggler.in/internal/db/query"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
)

//Handler has http handler functions for user APIs
type Handler struct {
	log        *zap.Logger
	reader     *request.Reader
	jsonWriter *response.JSONWriter
	repository *Repository
}

//NewHandler creates a new instance of Handler
func NewHandler(log *zap.Logger, reader *request.Reader, jsonWriter *response.JSONWriter, repository *Repository) *Handler {
	return &Handler{log: log, reader: reader, jsonWriter: jsonWriter, repository: repository}
}

func (h *Handler) addUser() http.HandlerFunc {

	type Request struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8"`
	}

	type Response struct {
		ID 	 	int32 `json:"id"`
		Name 	string `json:"name"`
		Email string `json:"email"`
		EmailVerified bool `json:"emailVerified"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		ok := h.reader.ReadJSONAndValidate(w, r, req)

		if !ok {
			h.log.Error("Error while validating fields")
			h.jsonWriter.Internal(w, r, &InternalError{})
			return
		}

		// Hashing password
		p := []byte(req.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)

		if err != nil {
			h.log.Error("Error while hashing password", zap.Error(err))
			h.jsonWriter.Internal(w, r, &InternalError{})
			return
		}

		user, err := h.repository.AddUser(r.Context(), query.AddUserParams{
			Name: req.Name,
			Email: req.Email,
			Password: string(hashedPassword),
		})

		if err != nil {
			h.log.Error("Failed adding user to DB", zap.Error(err))
			h.jsonWriter.DefaultError(w, r)
			return
		}

		// Send verification mail from here
		// Temporarily setting email as verified
		user, err = h.repository.VerifyEmail(r.Context(), user.ID)

		if err != nil {
			h.log.Error("Failed verifying email", zap.Error(err))
			h.jsonWriter.DefaultError(w, r)
			return
		}

		h.jsonWriter.Ok(w, r, &Response{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			EmailVerified: user.EmailVerified,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
}

func (h *Handler) getUser() http.HandlerFunc {
	type Response struct {
		Msg string `json:"msg"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		h.jsonWriter.Ok(w, r, &Response{
			Msg: "ok",
		})
	}
}