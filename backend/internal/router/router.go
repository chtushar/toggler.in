package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"

	"toggler.in/internal/auth"
	"toggler.in/internal/db"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
	"toggler.in/internal/users"
	"toggler.in/internal/validator"
)

type Config struct {
	R *mux.Router
	DB 	 *db.DB
	Log 	 *zap.Logger
	CS 	 *sessions.CookieStore
	JWTSecret string
}

func Routes(cfg *Config) {
	// Validator instance
	v := validator.New(cfg.Log)
	// JSON writer instance
	jw := response.NewJSONWriter(cfg.Log)
	// Request reader instance
	reader := request.NewReader(cfg.Log, jw, v)

	// User routes and handler
	ur := users.NewRepository(cfg.DB, cfg.Log)
	uh := users.NewHandler(cfg.Log, reader, jw, ur)
	users.UserRoutes(cfg.R, uh)

	// Auth routes and handler
	ar := auth.NewRepository(cfg.DB, cfg.Log)
	ah := auth.NewHandler(cfg.Log, reader, jw, ar, cfg.CS, cfg.JWTSecret)
	auth.AuthRoutes(cfg.R, ah)

}
