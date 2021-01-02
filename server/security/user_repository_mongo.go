package security

import (
	"context"
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var userRepositoryOnce sync.Once
var userRepositoryInstance *MongoUserRepository

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func ProvideMongoUserRepository(dbClient *mongo.Client) *MongoUserRepository {
	userRepositoryOnce.Do(func() {
		userRepositoryInstance = &MongoUserRepository{Collection: dbClient.Database("tasquest").Collection("users")}
	})
	return userRepositoryInstance
}

func (ur *MongoUserRepository) Save(user User) (User, error) {
	insertResult, err := ur.Collection.InsertOne(context.TODO(), user)

	if err != nil {
		return User{}, errors.Wrap(ErrFailedToSaveUser, err.Error())
	}

	insertedID := insertResult.InsertedID.(primitive.ObjectID).String()

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
	result := ur.Collection.FindOne(context.Background(), filter)
	user := User{}
	err := result.Decode(&user)

	if err != nil {
		return User{}, errors.Wrap(ErrUserNotFound, err.Error())
	}

	return user, nil
}
