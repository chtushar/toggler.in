package teams

import (
	"net/http"
	"time"

	"go.uber.org/zap"
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

func NewHandler(log *zap.Logger, reader *request.Reader, jsonWriter *response.JSONWriter, repository *Repository) *Handler {
	return &Handler{log: log, reader: reader, jsonWriter: jsonWriter, repository: repository}
}

func (h *Handler) createTeam() http.HandlerFunc  {
	type Request struct {
		Name string `json:"name" validate:"required"`
		OwnerID int32 `json:"owner_id" validate:"required"`
	}

	type Response struct {
		ID int32 `json:"id"`
		Name string `json:"name" validate:"required"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	return func (w http.ResponseWriter, r *http.Request)  {
		req := &Request{}
		ok := h.reader.ReadJSONAndValidate(w, r, req)

		if !ok {
			h.log.Error("Error while validating data")
			h.jsonWriter.Internal(w, r, &InternalError{})
			return
		}

		team, err := h.repository.CreateTeam(r.Context(), query.CreateTeamParams{
			Name: req.Name,
			OwnerID: req.OwnerID,
		})

		if err != nil {
			h.log.Error("Failed adding team to DB", zap.Error(err))
			h.jsonWriter.DefaultError(w, r)
			return
		}

		h.jsonWriter.Ok(w, r, &Response{
			ID: team.ID,
			Name: team.Name,
			CreatedAt: team.CreatedAt,
			UpdatedAt: team.UpdatedAt,
		})
	}
}