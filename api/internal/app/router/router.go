package router

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
	"toggler.in/internal/users"
	"toggler.in/internal/validator"
)

func Routes(r *mux.Router, db *gorm.DB, log *zap.Logger,) {
	// Validator instance
	v := validator.New(log)
	// JSON writer instance
	jw := response.NewJSONWriter(log)
	// Request reader instance
	reader := request.NewReader(log, jw, v)

	// User routes and handler
	uh := users.NewHandler(log, reader, jw, db)
	users.UserRoutes(r, uh)


}
