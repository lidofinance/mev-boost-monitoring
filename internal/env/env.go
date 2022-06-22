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
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.SetConfigFile(".env")

		// viper.AutomaticEnv()
		if viperErr := viper.ReadInConfig(); err != nil {
			if _, ok := viperErr.(viper.ConfigFileNotFoundError); !ok {
				err = viperErr
				return
			}
		}

		cfg = Config{
			AppConfig: AppConfig{
				Name:      viper.GetString("APP_NAME"),
				Env:       viper.GetString("ENV"),
				Port:      viper.GetUint("PORT"),
				LogFormat: viper.GetString("LOG_FORMAT"),
				LogLevel:  viper.GetString("LOG_LEVEL"),
				SentryDsn: viper.GetString("SENTRY_DSN"),
			},
			PgConfig: PgConfig{
				Port:     viper.GetUint("PG_PORT"),
				Host:     viper.GetString("PG_HOST"),
				Username: viper.GetString("PG_USERNAME"),
				Password: viper.GetString("PG_PASSWORD"),
				Database: viper.GetString("PG_DATABASE"),
				Schema:   viper.GetString("PG_SCHEMA"),
				SslMode:  viper.GetString("PG_SSLMODE"),
			},
		}
	})

	return &cfg, err
}
