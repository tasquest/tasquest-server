package adventurers

import (
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"tasquest.com/server/mongoutils"
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
	entity, err := mongoutils.Save(repo.collection, adventurer, Adventurer{})

	if err != nil {
		return Adventurer{}, errors.Wrap(ErrFailedToSaveAdventurer, err.Error())
	}

	return entity.(Adventurer), nil
}

func (repo *MongoAdventurerRepository) Update(adventurer Adventurer) (Adventurer, error) {
	updated, err := mongoutils.Update(repo.collection, adventurer.ID, adventurer, Adventurer{})

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	return updated.(Adventurer), nil
}

func (repo *MongoAdventurerRepository) FindByID(id string) (Adventurer, error) {
	entity, err := mongoutils.FindByID(repo.collection, id, Adventurer{})
	return entity.(Adventurer), err
}

func (repo *MongoAdventurerRepository) FindByUser(userID string) (Adventurer, error) {
	objectID, _ := primitive.ObjectIDFromHex(userID)
	return repo.FindByFilter(bson.M{"user_id": objectID})
}

func (repo *MongoAdventurerRepository) FindByFilter(filter bson.M) (Adventurer, error) {
	entity, err := mongoutils.FindByFilter(repo.collection, filter, Adventurer{})

	if err != nil {
		return Adventurer{}, errors.Wrap(ErrFailedToFetchAdventurer, err.Error())
	}

	return entity.(Adventurer), nil
}
