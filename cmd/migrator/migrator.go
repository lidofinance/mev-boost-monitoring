package main

import (
	"context"
	"os"

	"github.com/golang-migrate/migrate/v4"
	pgMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lidofinance/mev-boost-monitoring/internal/connectors/postgres"
	"github.com/lidofinance/mev-boost-monitoring/internal/env"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, envErr := env.Read(ctx)
	if envErr != nil {
		println("Read env error:", envErr.Error())

		os.Exit(1)
	}

	db, errDB := postgres.Connect(cfg.PgConfig)
	if errDB != nil {
		println("Connect db error:", errDB.Error())
		os.Exit(1)
	}

	driver, driverErr := pgMigrate.WithInstance(db.DB, &pgMigrate.Config{})
	if driverErr != nil {
		println("Migrate driver error:", driverErr.Error())
		os.Exit(1)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		cfg.PgConfig.Database, driver)
	if err != nil {
		println("Could not make migrate:", err.Error())
		os.Exit(1)
	}

	if err := m.Up(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
