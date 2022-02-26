package db

import (
	"fmt"
	"net/url"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"toggler.in/internal/configs"
)

var (
	once   sync.Once
	dbConn *gorm.DB
)

// Creates a new connection to the database.
func NewConnection(cfg *configs.App) (*gorm.DB, error) {
	return connect(cfg)
}

// Creates a new connection to the database if not present or returns the
// existing connection.
func GetConnection(cfg *configs.App) (*gorm.DB, error) {

	var (
		conn *gorm.DB
		err  error
	)

	once.Do(func() {
		conn, err = NewConnection(cfg)
		fmt.Println("Connected to database")
	})

	dbConn = conn

	return dbConn, err
}

// Connecting to the database.
func connect(cfg *configs.App) (*gorm.DB, error) {
	dsn := url.URL{
		User:   url.UserPassword(cfg.DB.User, cfg.DB.Password),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		Path:   cfg.DB.Name,
	}
	database, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	return database, nil
}
