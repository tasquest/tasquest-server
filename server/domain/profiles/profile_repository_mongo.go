package profiles

import (
	"context"
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var profileRepositoryOnce sync.Once
var profileRepositoryInstance *MongoProfileRepository

type MongoProfileRepository struct {
	collection *mongo.Collection
}

func ProvideMongoProfileRepository(dbClient *mongo.Client) *MongoProfileRepository {
	profileRepositoryOnce.Do(func() {
		profileRepositoryInstance = &MongoProfileRepository{
			collection: dbClient.Database("tasquest").Collection("profiles"),
		}
	})
	return profileRepositoryInstance
}

func (pr *MongoProfileRepository) Save(profile Profile) (Profile, error) {
	insertResult, err := pr.collection.InsertOne(context.Background(), profile)

	if err != nil {
		return Profile{}, errors.Wrap(ErrFailedToSaveProfile, err.Error())
	}

	insertedID := insertResult.InsertedID.(primitive.ObjectID).String()

	return pr.FindByID(insertedID)
}

func (pr *MongoProfileRepository) FindByID(id string) (Profile, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return pr.FindByFilter(bson.M{"_id": objectID})
}

func (pr *MongoProfileRepository) FindByUser(userID string) (Profile, error) {
	objectID, _ := primitive.ObjectIDFromHex(userID)
	return pr.FindByFilter(bson.M{"user_id": objectID})
}

func (pr *MongoProfileRepository) FindByFilter(filter bson.M) (Profile, error) {
	result := pr.collection.FindOne(context.Background(), filter)
	profile := Profile{}
	err := result.Decode(&profile)

	if err != nil {
		return profile, errors.Wrap(ErrFailedToFetchProfile, err.Error())
	}

	return profile, nil
}
