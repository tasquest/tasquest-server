package security

import (
	"emperror.dev/errors"
	"github.com/google/wire"
	"sync"
	"tasquest-server/src/security/commands"
	"tasquest-server/src/security/dao"
	"tasquest-server/src/security/models"
	securityError "tasquest-server/src/security/secerrors"
)

var SecurityProvider = wire.NewSet(ProvideSecurity, dao.UserRepositoryProvider)

type Security struct {
	userRepository *dao.UserRepository
}

var once sync.Once
var instance *Security

func ProvideSecurity(repository *dao.UserRepository) *Security {
	once.Do(func() {
		instance = &Security{userRepository: repository}
	})
	return instance
}

func (sec *Security) RegisterUser(command commands.RegisterUserCommand) (*models.User, error) {
	usr, err := sec.userRepository.FindByEmail(command.Email)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if usr != nil {
		return nil, errors.WithStack(securityError.ErrUserAlreadyExists)
	}

	newUser := models.User{
		Name:     command.Name,
		Surname:  command.Surname,
		Email:    command.Email,
		Password: command.Password,
	}

	return sec.userRepository.Save(newUser)
}

func (sec *Security) FetchUser(id string) (*models.User, error) {
	usr, err := sec.userRepository.FindByID(id)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return usr, nil
}
