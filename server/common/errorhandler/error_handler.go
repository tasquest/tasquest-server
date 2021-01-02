package errorhandler

import (
	"emperror.dev/emperror"
	logrushandler "emperror.dev/handler/logrus"
	"github.com/sirupsen/logrus"
	"sync"
)

var once sync.Once
var instance emperror.ErrorHandler

func ProvideErrorHandler(logger *logrus.Logger) emperror.ErrorHandler {
	once.Do(func() {
		handler := logrushandler.New(logger)
		instance = handler
	})
	return instance
}
