package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/zlog"
)

func RunMigrations(db *dbpg.DB, migrationsPath string) error {
	driver, err := postgres.WithInstance(db.Master, &postgres.Config{})
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to create migration driver")
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to create migrate instance")
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		zlog.Logger.Error().Err(err).Msg("failed to apply migrations")
		return err
	}

	zlog.Logger.Info().Msg("migrations applied successfully")
	return nil
}
