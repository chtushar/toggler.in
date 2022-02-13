package db

import (
	"fmt"
	"net/url"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"toggler.in/internal/configs"
)

var DB *gorm.DB

// Creates a new connection to the database.
func NewConnection(cfg *configs.App)  {
	connect(cfg)
}


var (
	once   sync.Once
)

// Creates a new connection to the database if not present or returns the
// existing connection.
func GetConnection(cfg *configs.App)  {
	once.Do(func() {
		NewConnection(cfg)
	})
}

// Connecting to the database.
func connect(cfg *configs.App)  {
	dsn := url.URL{
		User: url.UserPassword(cfg.DB.User, cfg.DB.Password),
		Scheme: "postgres",
		Host: fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		Path: cfg.DB.Name,
	}
	database, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = database
}

