package mongorepositories

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"tasquest.com/server/application/gamification/tasks"
	"tasquest.com/server/commons"
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

func (repo *MongoTaskRepository) Save(task tasks.Task) (tasks.Task, error) {
	savedTask, err := repo.collection.InsertOne(context.Background(), task)

	if err != nil {
		return tasks.Task{}, errors.WithStack(err)
	}

	insertedID := savedTask.InsertedID.(uuid.UUID)

	return repo.FindByID(insertedID)
}

func (repo *MongoTaskRepository) Update(task tasks.Task) (tasks.Task, error) {
	insertResult, err := repo.collection.UpdateOne(context.Background(), bson.M{"_id": task.ID}, task)

	if err != nil {
		return tasks.Task{}, errors.WithStack(err)
	}

	insertedID := insertResult.UpsertedID.(uuid.UUID)

	return repo.FindByID(insertedID)
}

func (repo *MongoTaskRepository) Delete(task tasks.Task) error {
	return repo.DeleteByID(task.ID)
}

func (repo *MongoTaskRepository) DeleteByID(id uuid.UUID) error {
	_, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": id})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *MongoTaskRepository) FindByID(id uuid.UUID) (tasks.Task, error) {
	return repo.FindByFilter(commons.Map{"_id": id})
}

func (repo *MongoTaskRepository) FindByFilter(filter commons.Map) (tasks.Task, error) {
	task := tasks.Task{}
	result := repo.collection.FindOne(context.Background(), filter)
	err := result.Decode(&task)

	if err != nil {
		return tasks.Task{}, errors.Wrap(tasks.ErrFailedToFetchTask, err.Error())
	}

	return task, nil
}
