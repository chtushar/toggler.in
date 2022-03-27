package db

import (
	"fmt"
	"net/url"
	"sync"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"toggler.in/internal/configs"
	"toggler.in/internal/models"
)

var (
	once   sync.Once
	dbConn *gorm.DB
)

// Creates a new connection to the database.
func NewConnection(cfg *configs.App, logger *zap.Logger) (*gorm.DB, error) {
	return connect(cfg, logger)
}

// Creates a new connection to the database if not present or returns the
// existing connection.
func GetConnection(cfg *configs.App, logger *zap.Logger) (*gorm.DB, error) {

	var (
		conn *gorm.DB
		err  error
	)

	once.Do(func() {
		conn, err = NewConnection(cfg, logger)
		fmt.Println("Connected to database")
	})

	dbConn = conn

	return dbConn, err
}

// Connecting to the database.
func connect(cfg *configs.App, logger *zap.Logger) (*gorm.DB, error) {
	dsn := url.URL{
		User:   url.UserPassword(cfg.DB.User, cfg.DB.Password),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		Path:   cfg.DB.Name,
	}
	database, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		logger.Error("Failed to connect to DB", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	// Migrations
	database.AutoMigrate(&models.User{})

	return database, nil
}
