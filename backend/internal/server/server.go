package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"toggler.in/internal/http/handlers"
	"toggler.in/internal/router"
)

type Config struct {
	Port int
	Logger *zap.Logger
}

type Server struct {
	server http.Server

	router *mux.Router

	logger *zap.Logger

	db *gorm.DB

	connClose chan int
}

func NewServer(cfg *Config, db *gorm.DB) *Server {
	r := mux.NewRouter().StrictSlash(true)

	return &Server{
		server: http.Server{
			Addr:    fmt.Sprintf("%s:%d", "", cfg.Port),
			Handler: r,
		},
		logger: cfg.Logger,
		router: r,
		db:     db,
		connClose: make(chan int, 1),
	}
}

func (s *Server) Listen() {
	s.setup()
	fmt.Println("Starting server on port", s.server.Addr)
	s.logger.Info("Starting server...", zap.String("address", s.server.Addr))
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Fatal("HTTP server error", zap.Error(err))
		fmt.Println("HTTP Server error", err)
	}
}

func (s *Server) WaitForShutdown() {
	<-s.connClose
}

func (s *Server) setup() {
	defer s.graceFullShutdown()

	apiRouter := s.router.PathPrefix("/").Subrouter()
	router.Routes(apiRouter, s.db, s.logger)

	// Add handlers and middlewares below this line
	s.server.Handler = handlers.RecoveryHandler(s.logger)(s.server.Handler)
}

func (s *Server) graceFullShutdown() {
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)

		sig := <-sigint
		s.logger.Info("OS terminate signal received", zap.String("signal", sig.String()))

		s.logger.Debug("Shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := s.server.Shutdown(ctx)
		if err != nil {
			s.logger.Error("Error shutting down server", zap.Error(err))
		}

		close(s.connClose)
	}()
}