package dao

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"tasquest-server/src/infra/database"
	"tasquest-server/src/security/models"
	"tasquest-server/src/security/secerrors"
)

var once sync.Once
var instance *UserRepository

var UserRepositoryProvider = wire.NewSet(ProvideUserRepository, database.ProvideDatasource)

type UserRepository struct {
	collection *mongo.Collection
}

func ProvideUserRepository(dbClient *mongo.Client) *UserRepository {
	once.Do(func() {
		collection := dbClient.Database("tasquest").Collection("users")
		instance = &UserRepository{collection: collection}
	})
	return instance
}

func (ur *UserRepository) Save(user models.User) (*models.User, error) {
	insertResult, err := ur.collection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, errors.Wrap(secerrors.ErrFailedToSaveUser, err.Error())
	}

	insertedID := insertResult.InsertedID.(primitive.ObjectID).String()

	return ur.FindByID(insertedID)
}

func (ur *UserRepository) FindByID(id string) (*models.User, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return ur.FindByFilter(bson.M{"_id": objectID})
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	return ur.FindByFilter(bson.M{"email": email})
}

func (ur *UserRepository) FindByFilter(filter bson.M) (*models.User, error) {
	result := ur.collection.FindOne(context.Background(), filter)
	user := models.User{}
	err := result.Decode(&user)

	if err != nil {
		return nil, errors.Wrap(secerrors.ErrUserNotFound, err.Error())
	}

	return &user, nil
}
