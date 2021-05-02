package mongorepositories

import (
	"context"
	"sync"

	"emperror.dev/errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"tasquest.com/server/application/security"
	"tasquest.com/server/commons"
)

var IsUserRepositoryInstanced sync.Once
var userRepositoryInstance *UserRepository

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(dbClient *mongo.Database) *UserRepository {
	IsUserRepositoryInstanced.Do(func() {
		userRepositoryInstance = &UserRepository{collection: dbClient.Collection("users")}
	})
	return userRepositoryInstance
}

func (ur *UserRepository) Save(user security.User) (security.User, error) {
	savedTask, err := ur.collection.InsertOne(context.Background(), user)

	if err != nil {
		return security.User{}, errors.WithStack(err)
	}

	insertedID := savedTask.InsertedID.(uuid.UUID)

	return ur.FindByID(insertedID)
}

func (ur *UserRepository) FindByID(id uuid.UUID) (security.User, error) {
	return ur.FindByFilter(commons.Map{"_id": id})
}

func (ur *UserRepository) FindByEmail(email string) (security.User, error) {
	return ur.FindByFilter(commons.Map{"email": email})
}

func (ur *UserRepository) FindByFilter(filter commons.Map) (security.User, error) {
	user := security.User{}
	result := ur.collection.FindOne(context.Background(), filter)
	err := result.Decode(&user)

	if err != nil {
		return security.User{}, errors.Wrap(security.ErrUserNotFound, err.Error())
	}

	return user, nil
}
