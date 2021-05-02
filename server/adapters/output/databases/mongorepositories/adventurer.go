package mongorepositories

import (
	"context"
	"sync"

	"emperror.dev/errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/commons"
)

var IsMongoAdventurerRepositoryInstanced sync.Once
var mongoAdventurerRepositoryInstance *AdventurerRepository

type AdventurerRepository struct {
	collection *mongo.Collection
}

func NewAdventurerRepository(dbClient *mongo.Database) *AdventurerRepository {
	IsMongoAdventurerRepositoryInstanced.Do(func() {
		mongoAdventurerRepositoryInstance = &AdventurerRepository{
			collection: dbClient.Collection("adventurers"),
		}
	})
	return mongoAdventurerRepositoryInstance
}

func (repo *AdventurerRepository) Save(adventurer adventurers.Adventurer) (adventurers.Adventurer, error) {
	insertResult, err := repo.collection.InsertOne(context.Background(), adventurer)

	if err != nil {
		return adventurers.Adventurer{}, errors.WithStack(err)
	}

	insertedID := insertResult.InsertedID.(uuid.UUID)

	return repo.FindByID(insertedID)
}

func (repo *AdventurerRepository) Update(adventurer adventurers.Adventurer) (adventurers.Adventurer, error) {
	insertResult, err := repo.collection.UpdateOne(context.Background(), bson.M{"_id": adventurer.ID}, adventurer)

	if err != nil {
		return adventurers.Adventurer{}, errors.WithStack(err)
	}

	insertedID := insertResult.UpsertedID.(uuid.UUID)

	return repo.FindByID(insertedID)
}

func (repo *AdventurerRepository) FindByID(id uuid.UUID) (adventurers.Adventurer, error) {
	return repo.FindByFilter(commons.Map{"_id": id})
}

func (repo *AdventurerRepository) FindByUser(userID uuid.UUID) (adventurers.Adventurer, error) {
	return repo.FindByFilter(commons.Map{"user_id": userID})
}

func (repo *AdventurerRepository) FindByFilter(filter commons.Map) (adventurers.Adventurer, error) {
	user := adventurers.Adventurer{}
	result := repo.collection.FindOne(context.Background(), filter)
	err := result.Decode(&user)

	if err != nil {
		return adventurers.Adventurer{}, errors.Wrap(adventurers.ErrFailedToFetchAdventurer, err.Error())
	}

	return user, nil
}
