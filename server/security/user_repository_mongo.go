package security

import (
	"context"
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var IsUserRepositoryInstanced sync.Once
var userRepositoryInstance *MongoUserRepository

type MongoUserRepository struct {
	collection *mongo.Collection
}

func ProvideMongoUserRepository(dbClient *mongo.Database) *MongoUserRepository {
	IsUserRepositoryInstanced.Do(func() {
		userRepositoryInstance = &MongoUserRepository{collection: dbClient.Collection("users")}
	})
	return userRepositoryInstance
}

func (ur *MongoUserRepository) Save(user User) (User, error) {
	savedTask, err := ur.collection.InsertOne(context.Background(), user)

	if err != nil {
		return User{}, errors.WithStack(err)
	}

	insertedID := savedTask.InsertedID.(primitive.ObjectID).String()

	return ur.FindByID(insertedID)
}

func (ur *MongoUserRepository) FindByID(id string) (User, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return ur.FindByFilter(bson.M{"_id": objectID})
}

func (ur *MongoUserRepository) FindByEmail(email string) (User, error) {
	return ur.FindByFilter(bson.M{"email": email})
}

func (ur *MongoUserRepository) FindByFilter(filter bson.M) (User, error) {
	user := User{}
	result := ur.collection.FindOne(context.Background(), filter)
	err := result.Decode(&user)

	if err != nil {
		return User{}, errors.Wrap(ErrUserNotFound, err.Error())
	}

	return user, nil
}
