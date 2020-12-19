package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
	"sync"
)

var once sync.Once
var instance *gin.Engine

func ProvideWebserver() *gin.Engine {
	once.Do(func() {
		logrus.Info("Configuring webserver")
		engine := gin.New()

		engine.Use(ginlogrus.Logger(logrus.New()))
		engine.Use(gin.Recovery())

		instance = engine
	})
	return instance
}
