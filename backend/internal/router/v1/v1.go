package v1

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"go.uber.org/zap"
	"toggler.in/internal/auth"
	"toggler.in/internal/db"
	"toggler.in/internal/http/request"
	"toggler.in/internal/http/response"
	"toggler.in/internal/users"
	"toggler.in/internal/validator"
)

type Config struct {
	R 	 *mux.Router
	DB 	 *db.DB
	Log  *zap.Logger
	SC 	 *securecookie.SecureCookie
	JWTSecret string
}

func V1Route(cfg *Config)  {

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
	ah := auth.NewHandler(&auth.Config{
		Log: cfg.Log,
		Reader: reader,
		JSONWriter: jw,
		Repository: ar,
		SecureCookie: cfg.SC,
		JWTSecret: cfg.JWTSecret,
	})
	auth.AuthRoutes(cfg.R, ah)
}