package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/zap"
)

// migrationVersion defines the current migration version. This ensures the app
// is always compatible with the version of the database.
const migrationVersion = 1

//go:embed migrations/*.sql
var migrations embed.FS

//buildMigrationClient get the new migrate instance
//source & target connection are needed to close it in the calling function
func buildMigrationClient(db *sql.DB) (*migrate.Migrate, error) {
	// Read the source for the migrations.
	// Our source is the SQL files in the migrations folder
	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to read migration source %w", err)
	}

	// Connect with the target i.e, our postgres DB
	target, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return nil, fmt.Errorf("failed to read migration target %w", err)
	}

	// Create a new instance of the migration using the defined source and target
	m, err := migrate.NewWithInstance("iofs", source, "postgres", target)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration instance %w", err)
	}

	return m, nil
}

//upMigrations migrates the postgres schema to the current version
func runUpMigrations(db *sql.DB, logger *zap.Logger) error {
	m, err := buildMigrationClient(db)
	if err != nil {
		logger.Error("failed to build client for up migration", zap.Error(err))
		return err
	}

	// Migrate the DB to the current version
	err = m.Migrate(migrationVersion)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		//If the error is present, and the error is not the No change error
		// log the error
		logger.Error("failed to run up migration", zap.Error(err), zap.Int("version", migrationVersion))
		return err
	}

	defer func(m *migrate.Migrate) {
		sErr, tErr := m.Close()
		if sErr != nil {
			logger.Error("Failed to close source", zap.Error(err))
		}

		if tErr != nil {
			logger.Error("Failed to close target", zap.Error(err))
		}
	}(m)

	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		logger.Info("Database is already migrated. Nothing to change", zap.Int("version", migrationVersion))
		return nil
	}

	logger.Info("Successfully executed migrations for DB", zap.Int("version", migrationVersion))
	return err
}

//runDownMigration migrates the postgres schema from the active migration version to all the way down
func runDownMigration(db *sql.DB, logger *zap.Logger) error {
	m, err := buildMigrationClient(db)
	if err != nil {
		logger.Error("failed to build client for down migration", zap.Error(err))
		return err
	}

	//applying all down migrations
	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		//If the error is present, and the error is not the No change error
		// log the error
		logger.Error("failed to run migration", zap.Error(err), zap.Int("version", migrationVersion))
		return err
	}

	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		logger.Info("Database is already down migrated. Nothing to change", zap.Int("version", migrationVersion))
		return nil
	}

	logger.Info("Database is now reset")

	return nil
}