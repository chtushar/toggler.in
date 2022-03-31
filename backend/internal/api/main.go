package app

import (
	"context"

	"toggler.in/internal/configs"
	"toggler.in/internal/db"
	"toggler.in/internal/logger"
	"toggler.in/internal/server"
)

func Execute() {
	cfg := configs.Get()
	log := logger.New(&logger.Config{Production: cfg.Production})

	// Connecting to the database.
	dbConn, err := db.GetConnection(context.Background(), cfg, log)

	if err != nil {
		panic(err)
	}

	// Initializing the server.
	srv := server.NewServer(&server.Config{Port: cfg.Port, Logger: log}, dbConn)
	srv.Listen()
}
