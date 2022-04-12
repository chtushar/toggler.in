package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
	"toggler.in/internal/configs"
)

type DB struct {
	*pgx.Conn
}

var (
	once   sync.Once
	dbConn *DB
)

//Setup runs the database migrations upto the current version
func Setup(cfg *configs.App, logger *zap.Logger) error {
	pgxConfig, err := getPgxConfig(cfg, logger)

	if err != nil {
		return err
	}

	err = runUpMigrations(stdlib.OpenDB(*pgxConfig), logger)
	if err != nil {
		logger.Error("failed to run migrations", zap.Error(err))
		return err
	}

	return nil
}

//Down to run the down migrations
func Down(ctx context.Context, cfg *configs.App, logger *zap.Logger) error {
	connConfig, err := getPgxConfig(cfg, logger)
	if err != nil {
		return err
	}

	_, err = pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		logger.Error("Failed to connect to DB", zap.Error(err))
		return fmt.Errorf("failed to connect to DB: %w", err)
	}

	//execute the .down.sql files
	return runDownMigration(stdlib.OpenDB(*connConfig), logger)
}

// Creates a new connection to the database.
func NewConnection(ctx context.Context,cfg *configs.App, logger *zap.Logger) (*DB, error) {
	return connect(ctx, cfg, logger)
}

// Creates a new connection to the database if not present or returns the
// existing connection.
func GetConnection(ctx context.Context,cfg *configs.App, logger *zap.Logger) (*DB, error) {

	var (
		conn *DB
		err  error
	)

	once.Do(func() {
		conn, err = NewConnection(ctx, cfg, logger)
		fmt.Println("Connected to database")
	})

	dbConn = conn

	return dbConn, err
}

//getPgxConfig builds and returns the pgx connection config
func getPgxConfig(cfg *configs.App, logger *zap.Logger) (*pgx.ConnConfig, error) {
	sslMode := "prefer"

	if cfg.DB.ForceTLS {
		sslMode = "require"
	}

	// postgres://username:password@localhost:5432/database_name?sslmode=disable
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		sslMode,
	)

	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		logger.Error("Failed to parse connection string", zap.Error(err))
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	connConfig.LogLevel = pgx.LogLevelDebug

	return connConfig, nil
}


// Connecting to the database.
func connect(ctx context.Context,cfg *configs.App, logger *zap.Logger) (*DB, error) {

	pgxConfig, err := getPgxConfig(cfg, logger)

	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(ctx, pgxConfig)


	if err != nil {
		logger.Error("Failed to connect to DB", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	return &DB{
		Conn: conn,
	}, nil
}
