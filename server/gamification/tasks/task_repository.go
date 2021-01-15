package tasks

import "go.mongodb.org/mongo-driver/bson"

type TaskRepository interface {
	Save(task Task) (Task, error)
	Update(task Task) (Task, error)
	Delete(task Task) (Task, error)
	DeleteByID(id string) (Task, error)
	FindByID(id string) (Task, error)
	FindByFilter(filter bson.M) (Task, error)
}
