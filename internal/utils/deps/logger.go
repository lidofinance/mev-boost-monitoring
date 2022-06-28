package deps

import "github.com/sirupsen/logrus"

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
}
