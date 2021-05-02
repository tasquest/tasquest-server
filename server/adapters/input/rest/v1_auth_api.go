package rest

import (
	"net/http"
	"sync"

	"emperror.dev/emperror"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"tasquest.com/server/application/security"
	"tasquest.com/server/commons"
)

type AuthAPI struct {
	handler  emperror.ErrorHandler
	router   *gin.Engine
	security security.UserService
}

var AuthAPIOnce sync.Once
var instance *AuthAPI

func NewAuthApi(router *gin.Engine, security security.UserService, handler emperror.ErrorHandler) *AuthAPI {
	AuthAPIOnce.Do(func() {
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
	var command security.RegisterUserCommand

	if err := c.ShouldBindJSON(&command); err != nil {
		auth.handler.Handle(err)
		c.JSON(http.StatusBadRequest, commons.ApplicationError{
			Title:   "Invalid Requisition",
			Message: err.Error(),
		})
		return
	}

	registered, err := auth.security.RegisterUser(command)

	if err != nil {
		auth.handler.Handle(err)
		appErr, _ := commons.ParseError(err)
		c.JSON(appErr.HTTPCode, appErr)
		return
	}

	c.JSON(http.StatusOK, registered)
}

func (auth *AuthAPI) fetchUser(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	usr, err := auth.security.FetchUser(id)

	if err != nil {
		auth.handler.Handle(err)
		appErr, _ := commons.ParseError(err)
		c.JSON(appErr.HTTPCode, appErr)
		return
	}

	c.JSON(http.StatusOK, usr)
}
