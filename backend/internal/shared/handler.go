package shared

import (
	"go.uber.org/zap"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
)

type Handler struct {
	log        *zap.Logger
	reader     *request.Reader
	jsonWriter *response.JSONWriter
}