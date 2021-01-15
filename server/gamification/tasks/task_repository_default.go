package tasks

import (
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"tasquest.com/server/mongoutils"
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
	savedTask, err := mongoutils.Save(repo.collection, task, Task{})
	return savedTask.(Task), err
}

func (repo *MongoTaskRepository) Update(task Task) (Task, error) {
	updatedTask, err := mongoutils.Update(repo.collection, task.ID, task, Task{})

	if err != nil {
		return Task{}, errors.WithStack(err)
	}

	return updatedTask.(Task), nil
}

func (repo *MongoTaskRepository) Delete(task Task) (Task, error) {
	_, err := mongoutils.Delete(repo.collection, task)

	if err != nil {
		return Task{}, errors.WithStack(err)
	}

	return task, nil
}

func (repo *MongoTaskRepository) DeleteByID(id string) (Task, error) {
	deletedTask, err := mongoutils.DeleteByID(repo.collection, id, Task{})
	return deletedTask.(Task), err
}

func (repo *MongoTaskRepository) FindByID(id string) (Task, error) {
	task, err := mongoutils.FindByID(repo.collection, id, Task{})
	return task.(Task), err
}

func (repo *MongoTaskRepository) FindByFilter(filter bson.M) (Task, error) {
	task, err := mongoutils.FindByFilter(repo.collection, filter, Task{})
	return task.(Task), err
}
