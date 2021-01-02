package api

import (
	"emperror.dev/emperror"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"tasquest.com/server/api"
	"tasquest.com/server/security"
	"tasquest.com/tests/mocks"
	"testing"
)

func TestRegisterUserSuccessfully(t *testing.T) {
	// Server Setup
	api.AuthAPIOnce = sync.Once{}
	server := gin.Default()
	userManagementMock := new(mocks.UserManagement)
	api.ProvideAuthAPI(server, userManagementMock, emperror.NewTestHandler())

	// Command
	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "TESTINGWITHAVALIDPASSWORD",
		PasswordConfirmation: "TESTINGWITHAVALIDPASSWORD",
	}

	registeredUser := security.User{
		ID:        primitive.ObjectID{},
		Email:     "test@test.com",
		Password:  "HASHEDPASSWORD",
		Active:    false,
		Providers: nil,
	}

	userManagementMock.On("RegisterUser", command).Return(registeredUser, nil)

	b, _ := json.Marshal(command)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(string(b)))
	server.ServeHTTP(w, req)

	const expectedJSON = `
				{
   					"active":false,
   					"email":"test@test.com",
   					"id":"000000000000000000000000",
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
	userManagementMock := new(mocks.UserManagement)
	api.ProvideAuthAPI(server, userManagementMock, emperror.NewTestHandler())

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
	userManagementMock := new(mocks.UserManagement)
	api.ProvideAuthAPI(server, userManagementMock, emperror.NewTestHandler())

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
	userManagementMock := new(mocks.UserManagement)
	api.ProvideAuthAPI(server, userManagementMock, emperror.NewTestHandler())

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
	userManagementMock := new(mocks.UserManagement)
	api.ProvideAuthAPI(server, userManagementMock, emperror.NewTestHandler())

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
