package security

import (
	"github.com/google/uuid"
	"tasquest.com/server/commons"
)

type UserService interface {
	FetchUser(id uuid.UUID) (User, error)
	RegisterUser(command RegisterUserCommand) (User, error)
}

type UserPersistence interface {
	Save(user User) (User, error)
}

type UserFinder interface {
	FindByID(id uuid.UUID) (User, error)
	FindByEmail(email string) (User, error)
	FindByFilter(filter commons.Map) (User, error)
}
