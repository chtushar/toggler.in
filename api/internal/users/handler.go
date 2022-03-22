package users

import (
	"net/http"
	"time"

	"go.uber.org/zap"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
	"toggler.in/internal/models"
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
		ID 	 	uint32 `json:"id"`
		Name 	string `json:"name"`
		Email string `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		ok := h.reader.ReadJSONAndValidate(w, r, req)

		if !ok {
			return
		}

		user, err := h.repository.AddUser(r.Context(), models.AddUserParams{
			Name: req.Name,
			Email: req.Email,
			Password: req.Password,
		})

		if err != nil {
			h.log.Error("Failed adding user to DB", zap.Error(err))
			h.jsonWriter.DefaultError(w, r)
			return
		}

		h.jsonWriter.Ok(w, r, &Response{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
}