package mongorepositories

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/commons"
)

var IsMongoAdventurerRepositoryInstanced sync.Once
var mongoAdventurerRepositoryInstance *MongoAdventurerRepository

type MongoAdventurerRepository struct {
	collection *mongo.Collection
}

func NewMongoAdventurerRepository(dbClient *mongo.Database) *MongoAdventurerRepository {
	IsMongoAdventurerRepositoryInstanced.Do(func() {
		mongoAdventurerRepositoryInstance = &MongoAdventurerRepository{
			collection: dbClient.Collection("adventurers"),
		}
	})
	return mongoAdventurerRepositoryInstance
}

func (repo *MongoAdventurerRepository) Save(adventurer adventurers.Adventurer) (adventurers.Adventurer, error) {
	insertResult, err := repo.collection.InsertOne(context.Background(), adventurer)

	if err != nil {
		return adventurers.Adventurer{}, errors.WithStack(err)
	}

	insertedID := insertResult.InsertedID.(uuid.UUID)

	return repo.FindByID(insertedID)
}

func (repo *MongoAdventurerRepository) Update(adventurer adventurers.Adventurer) (adventurers.Adventurer, error) {
	insertResult, err := repo.collection.UpdateOne(context.Background(), bson.M{"_id": adventurer.ID}, adventurer)

	if err != nil {
		return adventurers.Adventurer{}, errors.WithStack(err)
	}

	insertedID := insertResult.UpsertedID.(uuid.UUID)

	return repo.FindByID(insertedID)
}

func (repo *MongoAdventurerRepository) FindByID(id uuid.UUID) (adventurers.Adventurer, error) {
	return repo.FindByFilter(commons.Map{"_id": id})
}

func (repo *MongoAdventurerRepository) FindByUser(userID uuid.UUID) (adventurers.Adventurer, error) {
	return repo.FindByFilter(commons.Map{"user_id": userID})
}

func (repo *MongoAdventurerRepository) FindByFilter(filter commons.Map) (adventurers.Adventurer, error) {
	user := adventurers.Adventurer{}
	result := repo.collection.FindOne(context.Background(), filter)
	err := result.Decode(&user)

	if err != nil {
		return adventurers.Adventurer{}, errors.Wrap(adventurers.ErrFailedToFetchAdventurer, err.Error())
	}

	return user, nil
}
