package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"go.uber.org/zap"

	"toggler.in/internal/db"
	v1 "toggler.in/internal/router/v1"
)

type Config struct {
	R 	 *mux.Router
	DB 	 *db.DB
	Log  *zap.Logger
	SC 	 *securecookie.SecureCookie
	JWTSecret string
}

func Routes(cfg *Config) {
	v1.V1Route(&v1.Config{
		R: cfg.R.PathPrefix("/v1").Subrouter(),
		DB: cfg.DB,
		Log: cfg.Log,
		SC: cfg.SC,
		JWTSecret: cfg.JWTSecret,
	})
}
