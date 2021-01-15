package security

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tasquest.com/server/security"
	"tasquest.com/tests/mocks"
	"testing"
)

func TestRegisterUserSuccessfully(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepository)
	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "12345",
		PasswordConfirmation: "12345",
	}

	expectedUser := security.User{
		ID:        primitive.ObjectID{},
		Email:     "test@test.com",
		Password:  "BCRYPT",
		Active:    false,
		Providers: nil,
	}

	userRepositoryMock.On("FindByEmail", command.Email).Return(security.User{}, nil).Once()
	userRepositoryMock.On("Save", mock.Anything).Return(expectedUser, nil).Once()

	userManagement := security.DefaultUserManagement{userRepository: userRepositoryMock}

	_, _ = userManagement.RegisterUser(command)

	userRepositoryMock.AssertExpectations(t)
}

func TestUserAlreadyExists(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepository)
	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "12345",
		PasswordConfirmation: "12345",
	}

	existingUser := security.User{
		ID:        primitive.ObjectID{},
		Email:     "test@test.com",
		Password:  "BCRYPT",
		Active:    false,
		Providers: nil,
	}

	userRepositoryMock.On("FindByEmail", command.Email).Return(existingUser, nil).Once()

	userManagement := security.DefaultUserManagement{userRepository: userRepositoryMock}

	_, err := userManagement.RegisterUser(command)

	assert.NotNil(t, err)
	assert.Equal(t, security.ErrUserAlreadyExists.Message, err.Error())

	userRepositoryMock.AssertExpectations(t)
}

func TestPasswordsNotMatching(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepository)
	command := security.RegisterUserCommand{
		Email:                "test@test.com",
		Password:             "12345",
		PasswordConfirmation: "123451",
	}

	userRepositoryMock.On("FindByEmail", command.Email).Return(security.User{}, nil).Once()

	userManagement := security.DefaultUserManagement{userRepository: userRepositoryMock}

	_, err := userManagement.RegisterUser(command)

	assert.NotNil(t, err)
	assert.Equal(t, security.ErrPasswordNotMatch.Message, err.Error())

	userRepositoryMock.AssertExpectations(t)
}

func TestFetchUserSuccessfully(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepository)
	anyID := "anyId"

	userRepositoryMock.On("FindByID", anyID).Return(security.User{}, nil).Once()

	userManagement := security.DefaultUserManagement{userRepository: userRepositoryMock}

	usr, err := userManagement.FetchUser(anyID)

	assert.Nil(t, err)
	assert.Equal(t, usr, security.User{})

	userRepositoryMock.AssertExpectations(t)
}
