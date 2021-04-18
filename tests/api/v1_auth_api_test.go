package api

import (
	"emperror.dev/emperror"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"tasquest.com/server/api"
	"tasquest.com/server/application/security"
	"tasquest.com/tests/mocks"
	"testing"
)

func TestRegisterUserSuccessfully(t *testing.T) {
	// Server Setup
	api.AuthAPIOnce = sync.Once{}
	server := gin.Default()
	userService := new(mocks.UserService)
	api.ProvideAuthAPI(server, userService, emperror.NewTestHandler())
	expectedUserId := uuid.MustParse("b2e02796-a051-11eb-b9ad-b42e99f46078")

	// Command
	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "TESTINGWITHAVALIDPASSWORD",
		PasswordConfirmation: "TESTINGWITHAVALIDPASSWORD",
	}

	registeredUser := security.User{
		ID:        expectedUserId,
		Email:     "test@test.com",
		Password:  "HASHEDPASSWORD",
		Active:    false,
		Providers: nil,
	}

	userService.On("RegisterUser", command).Return(registeredUser, nil)

	b, _ := json.Marshal(command)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(string(b)))
	server.ServeHTTP(w, req)

	const expectedJSON = `
				{
   					"active":false,
   					"email":"test@test.com",
   					"id":"b2e02796-a051-11eb-b9ad-b42e99f46078",
   					"password":"HASHEDPASSWORD",
   					"providers":null
				}
	`

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestRegisterUserMissingField(t *testing.T) {
	// Server Setup
	api.AuthAPIOnce = sync.Once{}
	server := gin.Default()
	userService := new(mocks.UserService)
	api.ProvideAuthAPI(server, userService, emperror.NewTestHandler())

	// Command
	command := security.RegisterUserCommand{
		Email: "test@test.com",
	}

	b, _ := json.Marshal(command)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(string(b)))
	server.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestRegisterUserMissingFields(t *testing.T) {
	// Server Setup
	api.AuthAPIOnce = sync.Once{}
	server := gin.Default()
	userService := new(mocks.UserService)
	api.ProvideAuthAPI(server, userService, emperror.NewTestHandler())

	// Command
	command := security.RegisterUserCommand{}

	b, _ := json.Marshal(command)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(string(b)))
	server.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestRegisterUserSmallPassword(t *testing.T) {
	// Server Setup
	api.AuthAPIOnce = sync.Once{}
	server := gin.Default()
	userService := new(mocks.UserService)
	api.ProvideAuthAPI(server, userService, emperror.NewTestHandler())

	// Command
	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "small",
		PasswordConfirmation: "small",
	}

	b, _ := json.Marshal(command)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(string(b)))
	server.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestRegisterUserInvalidEmail(t *testing.T) {
	// Server Setup
	api.AuthAPIOnce = sync.Once{}
	server := gin.Default()
	userService := new(mocks.UserService)
	api.ProvideAuthAPI(server, userService, emperror.NewTestHandler())

	// Command
	command := security.RegisterUserCommand{
		Email:                "test@@test.com",
		Password:             "TESTINGWITHAVALIDPASSWORD",
		PasswordConfirmation: "TESTINGWITHAVALIDPASSWORD",
	}

	b, _ := json.Marshal(command)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(string(b)))
	server.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
