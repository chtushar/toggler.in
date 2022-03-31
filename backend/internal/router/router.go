package router

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"toggler.in/internal/db"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
	"toggler.in/internal/users"
	"toggler.in/internal/validator"
)

func Routes(r *mux.Router, db *db.DB, log *zap.Logger,) {
	// Validator instance
	v := validator.New(log)
	// JSON writer instance
	jw := response.NewJSONWriter(log)
	// Request reader instance
	reader := request.NewReader(log, jw, v)

	// User routes and handler
	ur := users.NewRepository(db, log)
	uh := users.NewHandler(log, reader, jw, ur)
	users.UserRoutes(r, uh)

}
