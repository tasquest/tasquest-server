package logs

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var once sync.Once
var instance *logrus.Logger

func ProvideLogger() *logrus.Logger {
	once.Do(func() {
		logger := logrus.New()
		instance = logger
	})
	return instance
}
