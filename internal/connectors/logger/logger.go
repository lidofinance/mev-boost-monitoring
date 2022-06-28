package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/lidofinance/mev-boost-monitoring/internal/env"
	"github.com/onrik/logrus/sentry"
)

var (
	logger            *logrus.Logger
	onceDefaultClient sync.Once
)

func New(cfg *env.AppConfig) (*logrus.Logger, error) {
	var (
		err error
	)

	onceDefaultClient.Do(func() {
		logger = logrus.StandardLogger()

		logLevel, levelErr := logrus.ParseLevel(cfg.LogLevel)
		if levelErr != nil {
			err = levelErr
			return
		}

		logger.SetLevel(logLevel)
		logger.SetFormatter(&logrus.TextFormatter{})

		if cfg.LogFormat == "json" {
			logger.SetFormatter(&logrus.JSONFormatter{})
		}

		logger.SetOutput(os.Stdout)

		if cfg.SentryDsn != "" {
			hook, sentryErr := sentry.NewHook(
				sentry.Options{
					Dsn:              cfg.SentryDsn,
					AttachStacktrace: true,
				},
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
				logrus.DebugLevel,
				logrus.InfoLevel,
			)

			if sentryErr != nil {
				err = sentryErr
				return
			}

			logger.Hooks.Add(hook)
		}
	})

	return logger, err
}
