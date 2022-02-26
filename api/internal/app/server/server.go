package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"toggler.in/internal/app/router"
)

type Config struct {
	Port int
}

type Server struct {
	server http.Server

	router *mux.Router

	db *gorm.DB
}

func NewServer(cfg *Config, db *gorm.DB) *Server {
	r := mux.NewRouter().StrictSlash(true)

	return &Server{
		server: http.Server{
			Addr:    fmt.Sprintf("%s:%d", "", cfg.Port),
			Handler: r,
		},
		router: r,
		db:     db,
	}
}

func (s *Server) Listen() {
	s.setup()
	fmt.Println("Starting server on port", s.server.Addr)
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		// TODO: Logger goes here
		fmt.Println("HTTP Server error", err)
	}
}

func (s *Server) setup() {
	apiRouter := s.router.PathPrefix("/").Subrouter()
	router.Routes(apiRouter, s.db)

	// Add handlers and middlewares below this line
}
