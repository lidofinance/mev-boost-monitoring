package postgres

import (
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4/stdlib" // justifying it
	"github.com/jmoiron/sqlx"

	"github.com/lidofinance/go-template/internal/env"
)

var (
	db                *sqlx.DB
	onceDefaultClient sync.Once
)

const (
	MaxOpenConns = 25
	MaxIdleConns = 60 * int(time.Second)
)

func Connect(config env.PgConfig) (*sqlx.DB, error) {

	conf, parseErr := pgx.ParseConfig(
		fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`,
			config.Host, config.Port, config.Username, config.Password, config.Schema, config.SslMode),
	)

	if parseErr != nil {
		return nil, parseErr
	}

	conf.PreferSimpleProtocol = true
	conf.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on",
	}

	onceDefaultClient.Do(func() {
		db = sqlx.NewDb(stdlib.OpenDB(*conf), "pgx")

		// force a connection and test that it worked
		err := db.Ping()
		if err != nil {
			return
		}

		db.SetMaxOpenConns(MaxOpenConns)
		db.SetMaxIdleConns(MaxIdleConns)
	})

	return db, nil
}
