package env

import (
	"context"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig AppConfig
	PgConfig  PgConfig
}

type AppConfig struct {
	Name      string
	Env       string
	URL       string
	Port      uint
	LogFormat string
	LogLevel  string
	SentryDsn string
}

type PgConfig struct {
	Port     uint
	Host     string
	Username string
	Password string
	Database string
	Schema   string
	SslMode  string
}

var (
	cfg Config

	onceDefaultClient sync.Once
)

func Read(ctx context.Context) (*Config, error) {
	var err error

	onceDefaultClient.Do(func() {
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")

		if err = viper.ReadInConfig(); err != nil {
			return
		}

		cfg = Config{
			AppConfig: AppConfig{
				Name:      viper.GetString("app.name"),
				Env:       viper.GetString("app.env"),
				URL:       viper.GetString("app.url"),
				Port:      viper.GetUint("app.port"),
				LogFormat: viper.GetString("app.logFormat"),
				LogLevel:  viper.GetString("app.logLevel"),
				SentryDsn: viper.GetString("app.sentryDsn"),
			},
			PgConfig: PgConfig{
				Port:     viper.GetUint("pg.port"),
				Host:     viper.GetString("pg.host"),
				Username: viper.GetString("pg.username"),
				Password: viper.GetString("pg.password"),
				Database: viper.GetString("pg.database"),
				Schema:   viper.GetString("pg.schema"),
				SslMode:  viper.GetString("pg.sslmode"),
			},
		}
	})

	return &cfg, err
}
