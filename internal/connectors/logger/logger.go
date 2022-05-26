package logger

import (
	"os"
	"sync"

	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"

	"github.com/lidofinance/go-template/internal/env"
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
			hook, hookErr := logrus_sentry.NewSentryHook(cfg.SentryDsn, []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
			})

			if hookErr != nil {
				err = hookErr
				return
			}

			logger.Hooks.Add(hook)
		}
	})

	return logger, err
}
