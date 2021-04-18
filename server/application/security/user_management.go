package security

import (
	"emperror.dev/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type UserManagement struct {
	userFinder      UserFinder
	userPersistence UserPersistence
}

var IsUserManagementInstanced sync.Once
var userManagementInstance *UserManagement

func NewUserManagement(finder UserFinder, saver UserPersistence) *UserManagement {
	IsUserManagementInstanced.Do(func() {
		userManagementInstance = &UserManagement{
			userFinder:      finder,
			userPersistence: saver,
		}
	})
	return userManagementInstance
}

func (um *UserManagement) RegisterUser(command RegisterUserCommand) (User, error) {
	existingUser, err := um.userFinder.FindByEmail(command.Email)

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

	return um.userPersistence.Save(newUser)
}

func (um *UserManagement) FetchUser(id uuid.UUID) (User, error) {
	usr, err := um.userFinder.FindByID(id)

	if err != nil {
		return User{}, errors.WithStack(err)
	}

	return usr, nil
}
