package cmd

import (
	"toggler.in/internal/configs"
	"toggler.in/internal/db"
)

func Execute()  {
	cfg := configs.Get()

	// Connecting to the database.
	db.GetConnection(cfg)

}