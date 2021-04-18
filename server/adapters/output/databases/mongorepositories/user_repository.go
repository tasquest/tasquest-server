package mongorepositories

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"tasquest.com/server/application/security"
	"tasquest.com/server/commons"
)

var IsUserRepositoryInstanced sync.Once
var userRepositoryInstance *MongoUserRepository

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(dbClient *mongo.Database) *MongoUserRepository {
	IsUserRepositoryInstanced.Do(func() {
		userRepositoryInstance = &MongoUserRepository{collection: dbClient.Collection("users")}
	})
	return userRepositoryInstance
}

func (ur *MongoUserRepository) Save(user security.User) (security.User, error) {
	savedTask, err := ur.collection.InsertOne(context.Background(), user)

	if err != nil {
		return security.User{}, errors.WithStack(err)
	}

	insertedID := savedTask.InsertedID.(uuid.UUID)

	return ur.FindByID(insertedID)
}

func (ur *MongoUserRepository) FindByID(id uuid.UUID) (security.User, error) {
	return ur.FindByFilter(commons.Map{"_id": id})
}

func (ur *MongoUserRepository) FindByEmail(email string) (security.User, error) {
	return ur.FindByFilter(commons.Map{"email": email})
}

func (ur *MongoUserRepository) FindByFilter(filter commons.Map) (security.User, error) {
	user := security.User{}
	result := ur.collection.FindOne(context.Background(), filter)
	err := result.Decode(&user)

	if err != nil {
		return security.User{}, errors.Wrap(security.ErrUserNotFound, err.Error())
	}

	return user, nil
}
