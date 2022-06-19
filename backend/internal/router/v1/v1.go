package v1

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"go.uber.org/zap"
	"toggler.in/internal/auth"
	"toggler.in/internal/db"
	"toggler.in/internal/helpers"
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
	JWT  *helpers.JWT
}

func V1Route(cfg *Config)  {

	// Validator instance
	v := validator.New(cfg.Log)
	// JSON writer instance
	jw := response.NewJSONWriter(cfg.Log)
	// Request reader instance
	reader := request.NewReader(cfg.Log, jw, v)

	// User routes and handler
	usersRoute(cfg, reader, jw)

	// Auth routes and handler
	authRoutes(cfg, reader, jw)
}

func usersRoute(cfg *Config, reader *request.Reader, jw *response.JSONWriter) {
	ur := users.NewRepository(cfg.DB, cfg.Log)
	uh := users.NewHandler(cfg.Log, reader, jw, ur)
	users.UserRoutes(cfg.R, uh)
}

func authRoutes(cfg *Config, reader *request.Reader, jw *response.JSONWriter) {
	ar := auth.NewRepository(cfg.DB, cfg.Log)
	ah := auth.NewHandler(&auth.Config{
		Log: cfg.Log,
		Reader: reader,
		JSONWriter: jw,
		Repository: ar,
		SecureCookie: cfg.SC,
		JWT: cfg.JWT,
	})
	auth.AuthRoutes(cfg.R, ah)
}
