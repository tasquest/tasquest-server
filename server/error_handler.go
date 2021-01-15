package server

import (
	"emperror.dev/emperror"
	logrushandler "emperror.dev/handler/logrus"
	"github.com/sirupsen/logrus"
	"sync"
)

var IsErrorHandlerInstanced sync.Once
var errorHandlerInstance emperror.ErrorHandler

func ProvideErrorHandler(logger *logrus.Logger) emperror.ErrorHandler {
	IsErrorHandlerInstanced.Do(func() {
		handler := logrushandler.New(logger)
		errorHandlerInstance = handler
	})
	return errorHandlerInstance
}
