package users

import (
	"net/http"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
)

//Handler has http handler functions for user APIs
type Handler struct {
	log        *zap.Logger
	reader     *request.Reader
	jsonWriter *response.JSONWriter
	db *gorm.DB
}

//NewHandler creates a new instance of Handler
func NewHandler(log *zap.Logger, reader *request.Reader, jsonWriter *response.JSONWriter, db *gorm.DB) *Handler {
	return &Handler{log: log, reader: reader, jsonWriter: jsonWriter, db: db}
}

func (h *Handler) addUser() http.HandlerFunc {

	type Response struct {
		Message string `json:"message"`
	}

	resp := Response{
		Message: "User added successfully",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		h.jsonWriter.Ok(w, r, resp)
	}
}