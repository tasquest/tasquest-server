package v1

import (
	"emperror.dev/emperror"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"tasquest-server/src/common/errorhandler"
	"tasquest-server/src/security"
	"tasquest-server/src/security/commands"
)

type AuthAPI struct {
	handler  emperror.ErrorHandler
	router   *gin.Engine
	security *security.Security
}

var once sync.Once
var instance *AuthAPI

func ProvideAuthAPI(router *gin.Engine, security *security.Security, handler emperror.ErrorHandler) *AuthAPI {
	once.Do(func() {
		instance = &AuthAPI{router: router, security: security, handler: handler}
		instance.registerApis()
	})
	return instance
}

func (auth *AuthAPI) registerApis() {
	auth.router.POST("/api/v1/user", auth.registerUser)
	auth.router.GET("/api/v1/user/:id", auth.fetchUser)
}

func (auth *AuthAPI) registerUser(c *gin.Context) {
	var command commands.RegisterUserCommand

	if err := c.ShouldBindJSON(&command); err != nil {
		auth.handler.Handle(err)
		c.JSON(http.StatusBadRequest, errorhandler.ApplicationError{
			Title:   "Invalid Requisition",
			Message: "Invalid Requisition",
		})
		return
	}

	registered, err := auth.security.RegisterUser(command)

	if err != nil {
		auth.handler.Handle(err)
		appErr, _ := errorhandler.ParseError(err)
		c.JSON(appErr.HTTPCode, appErr)
		return
	}

	c.JSON(http.StatusOK, registered)
}

func (auth *AuthAPI) fetchUser(c *gin.Context) {
	id := c.Param("id")
	usr, err := auth.security.FetchUser(id)

	if err != nil {
		auth.handler.Handle(err)
		appErr, _ := errorhandler.ParseError(err)
		c.JSON(appErr.HTTPCode, appErr)
		return
	}

	c.JSON(http.StatusOK, usr)
}
