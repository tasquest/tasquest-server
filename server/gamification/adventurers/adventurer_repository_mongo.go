package adventurers

import (
	"context"
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var IsMongoAdventurerRepositoryInstanced sync.Once
var mongoAdventurerRepositoryInstance *MongoAdventurerRepository

type MongoAdventurerRepository struct {
	collection *mongo.Collection
}

func ProvideMongoAdventurerRepository(dbClient *mongo.Database) *MongoAdventurerRepository {
	IsMongoAdventurerRepositoryInstanced.Do(func() {
		mongoAdventurerRepositoryInstance = &MongoAdventurerRepository{
			collection: dbClient.Collection("adventurers"),
		}
	})
	return mongoAdventurerRepositoryInstance
}

func (repo *MongoAdventurerRepository) Save(adventurer Adventurer) (Adventurer, error) {
	insertResult, err := repo.collection.InsertOne(context.Background(), adventurer)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	insertedID := insertResult.InsertedID.(primitive.ObjectID).String()

	return repo.FindByID(insertedID)
}

func (repo *MongoAdventurerRepository) Update(adventurer Adventurer) (Adventurer, error) {
	insertResult, err := repo.collection.UpdateOne(context.Background(), bson.M{"_id": adventurer.ID}, adventurer)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	insertedID := insertResult.UpsertedID.(primitive.ObjectID).String()

	return repo.FindByID(insertedID)
}

func (repo *MongoAdventurerRepository) FindByID(id string) (Adventurer, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return repo.FindByFilter(bson.M{"_id": objectID})
}

func (repo *MongoAdventurerRepository) FindByUser(userID string) (Adventurer, error) {
	objectID, _ := primitive.ObjectIDFromHex(userID)
	return repo.FindByFilter(bson.M{"user_id": objectID})
}

func (repo *MongoAdventurerRepository) FindByFilter(filter bson.M) (Adventurer, error) {
	user := Adventurer{}
	result := repo.collection.FindOne(context.Background(), filter)
	err := result.Decode(&user)

	if err != nil {
		return Adventurer{}, errors.Wrap(ErrFailedToFetchAdventurer, err.Error())
	}

	return user, nil
}
