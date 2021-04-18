package security

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"tasquest.com/server/application/security"
	"tasquest.com/tests/mocks"
	"testing"
)

func TestRegisterUserSuccessfully(t *testing.T) {
	security.IsUserManagementInstanced = sync.Once{}
	userFinder := new(mocks.UserFinder)
	userPersistence := new(mocks.UserPersistence)
	userManagement := security.NewUserManagement(userFinder, userPersistence)
	expectedUserId, _ := uuid.NewUUID()

	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "12345",
		PasswordConfirmation: "12345",
	}

	expectedUser := security.User{
		ID:        expectedUserId,
		Email:     "test@test.com",
		Password:  "BCRYPT",
		Active:    false,
		Providers: nil,
	}

	userFinder.On("FindByEmail", command.Email).Return(security.User{}, nil).Once()
	userPersistence.On("Save", mock.Anything).Return(expectedUser, nil).Once()

	_, _ = userManagement.RegisterUser(command)

	userFinder.AssertExpectations(t)
	userPersistence.AssertExpectations(t)
}

func TestUserAlreadyExists(t *testing.T) {
	security.IsUserManagementInstanced = sync.Once{}
	userFinder := new(mocks.UserFinder)
	userPersistence := new(mocks.UserPersistence)
	userManagement := security.NewUserManagement(userFinder, userPersistence)
	existingUserId, _ := uuid.NewUUID()

	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "12345",
		PasswordConfirmation: "12345",
	}

	existingUser := security.User{
		ID:        existingUserId,
		Email:     "test@test.com",
		Password:  "BCRYPT",
		Active:    false,
		Providers: nil,
	}

	userFinder.On("FindByEmail", command.Email).Return(existingUser, nil).Once()

	_, err := userManagement.RegisterUser(command)

	assert.NotNil(t, err)
	assert.Equal(t, security.ErrUserAlreadyExists.Message, err.Error())

	userFinder.AssertExpectations(t)
	userPersistence.AssertExpectations(t)
}

func TestPasswordsNotMatching(t *testing.T) {
	security.IsUserManagementInstanced = sync.Once{}
	userFinder := new(mocks.UserFinder)
	userPersistence := new(mocks.UserPersistence)
	userManagement := security.NewUserManagement(userFinder, userPersistence)

	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "12345",
		PasswordConfirmation: "123451",
	}

	userFinder.On("FindByEmail", command.Email).Return(security.User{}, nil).Once()

	_, err := userManagement.RegisterUser(command)

	assert.NotNil(t, err)
	assert.Equal(t, security.ErrPasswordNotMatch.Message, err.Error())

	userFinder.AssertExpectations(t)
	userPersistence.AssertExpectations(t)
}

func TestFetchUserSuccessfully(t *testing.T) {
	security.IsUserManagementInstanced = sync.Once{}
	userFinder := new(mocks.UserFinder)
	userPersistence := new(mocks.UserPersistence)
	userManagement := security.NewUserManagement(userFinder, userPersistence)
	userId, _ := uuid.NewUUID()

	userFinder.On("FindByID", userId).Return(security.User{}, nil).Once()

	usr, err := userManagement.FetchUser(userId)

	assert.Nil(t, err)
	assert.Equal(t, usr, security.User{})

	userFinder.AssertExpectations(t)
	userPersistence.AssertExpectations(t)
}
