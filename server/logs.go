package server

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var IsLoggerInstanced sync.Once
var loggerInstance *logrus.Logger

func ProvideLogger() *logrus.Logger {
	IsLoggerInstanced.Do(func() {
		logger := logrus.New()
		loggerInstance = logger
	})
	return loggerInstance
}
