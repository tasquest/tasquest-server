package security

import (
	"emperror.dev/errors"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type DefaultUserManagement struct {
	UserRepository UserRepository
}

var userManagementOnce sync.Once
var userManagementInstance *DefaultUserManagement

func ProvideDefaultUserManagement(repository UserRepository) *DefaultUserManagement {
	userManagementOnce.Do(func() {
		userManagementInstance = &DefaultUserManagement{UserRepository: repository}
	})
	return userManagementInstance
}

func (um *DefaultUserManagement) RegisterUser(command RegisterUserCommand) (User, error) {
	existingUser, err := um.UserRepository.FindByEmail(command.Email)

	if err != nil {
		return User{}, errors.WithStack(err)
	}

	if !cmp.Equal(existingUser, User{}) {
		return User{}, errors.WithStack(ErrUserAlreadyExists)
	}

	if command.Password != command.PasswordConfirmation {
		return User{}, errors.WithStack(ErrPasswordNotMatch)
	}

	hashedPassword, bcryptErr := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)

	if bcryptErr != nil {
		return User{}, errors.WithStack(ErrFailedPasswordGenerate)
	}

	newUser := User{
		Email:    command.Email,
		Password: string(hashedPassword),
		Active:   false,
	}

	return um.UserRepository.Save(newUser)
}

func (um *DefaultUserManagement) FetchUser(id string) (User, error) {
	usr, err := um.UserRepository.FindByID(id)

	if err != nil {
		return User{}, errors.WithStack(err)
	}

	return usr, nil
}
