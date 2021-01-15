package security

import (
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"tasquest.com/server/mongoutils"
)

var userRepositoryOnce sync.Once
var userRepositoryInstance *MongoUserRepository

type MongoUserRepository struct {
	collection *mongo.Collection
}

func ProvideMongoUserRepository(dbClient *mongo.Database) *MongoUserRepository {
	userRepositoryOnce.Do(func() {
		userRepositoryInstance = &MongoUserRepository{collection: dbClient.Collection("users")}
	})
	return userRepositoryInstance
}

func (ur *MongoUserRepository) Save(user User) (User, error) {
	insertResult, err := mongoutils.Save(ur.collection, user, User{})

	if err != nil {
		return User{}, errors.Wrap(ErrFailedToSaveUser, err.Error())
	}

	return insertResult.(User), nil
}

func (ur *MongoUserRepository) FindByID(id string) (User, error) {
	user, err := mongoutils.FindByID(ur.collection, id, User{})
	return user.(User), err
}

func (ur *MongoUserRepository) FindByEmail(email string) (User, error) {
	return ur.FindByFilter(bson.M{"email": email})
}

func (ur *MongoUserRepository) FindByFilter(filter bson.M) (User, error) {
	user, err := mongoutils.FindByFilter(ur.collection, filter, User{})

	if err != nil {
		return User{}, errors.Wrap(ErrUserNotFound, err.Error())
	}

	return user.(User), nil
}
