package cmd

import (
	"toggler.in/internal/app/server"
	"toggler.in/internal/configs"
	"toggler.in/internal/db"
)

func Execute() {
	cfg := configs.Get()

	// Connecting to the database.
	dbConn, err := db.GetConnection(cfg)

	if err != nil {
		panic(err)
	}

	// Initializing the server.
	srv := server.NewServer(&server.Config{Port: cfg.Port}, dbConn)
	srv.Listen()
}
