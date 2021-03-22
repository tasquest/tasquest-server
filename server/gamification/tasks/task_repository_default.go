package tasks

import (
	"context"
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type MongoTaskRepository struct {
	collection *mongo.Collection
}

var IsMongoTaskRepositoryInstanced sync.Once
var mongoTaskRepositoryInstance *MongoTaskRepository

func ProvideMongoTaskRepository(dbClient *mongo.Database) *MongoTaskRepository {
	IsMongoTaskRepositoryInstanced.Do(func() {
		mongoTaskRepositoryInstance = &MongoTaskRepository{
			collection: dbClient.Collection("tasks"),
		}
	})
	return mongoTaskRepositoryInstance
}

func (repo *MongoTaskRepository) Save(task Task) (Task, error) {
	savedTask, err := repo.collection.InsertOne(context.Background(), task)

	if err != nil {
		return Task{}, errors.WithStack(err)
	}

	insertedID := savedTask.InsertedID.(primitive.ObjectID).String()

	return repo.FindByID(insertedID)
}

func (repo *MongoTaskRepository) Update(task Task) (Task, error) {
	insertResult, err := repo.collection.UpdateOne(context.Background(), bson.M{"_id": task.ID}, task)

	if err != nil {
		return Task{}, errors.WithStack(err)
	}

	insertedID := insertResult.UpsertedID.(primitive.ObjectID).String()

	return repo.FindByID(insertedID)
}

func (repo *MongoTaskRepository) Delete(task Task) error {
	return repo.DeleteByID(task.ID.String())
}

func (repo *MongoTaskRepository) DeleteByID(id string) error {
	_, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": id})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *MongoTaskRepository) FindByID(id string) (Task, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return repo.FindByFilter(bson.M{"_id": objectID})
}

func (repo *MongoTaskRepository) FindByFilter(filter bson.M) (Task, error) {
	task := Task{}
	result := repo.collection.FindOne(context.Background(), filter)
	err := result.Decode(&task)

	if err != nil {
		return Task{}, errors.Wrap(ErrFailedToFetchTask, err.Error())
	}

	return task, nil
}
