package tasks

import (
	"github.com/google/uuid"

	"tasquest.com/server/commons"
)

type TaskService interface {
	CreateTask(command CreateTaskCommand) (Task, error)
	UpdateTask(taskID uuid.UUID, command UpdateTaskCommand) (Task, error)
	DeleteTask(taskID uuid.UUID) (Task, error)
	AdventurerCompletesTask(command AdventurerCompletedTaskCommand) error
}

type TaskPersistence interface {
	Save(task Task) (Task, error)
	Update(task Task) (Task, error)
	Delete(task Task) error
	DeleteByID(id uuid.UUID) error
}

type TaskFinder interface {
	FindByID(id uuid.UUID) (Task, error)
	FindOneByFilter(filter commons.SqlFilter) (Task, error)
	FindAllByFilter(filter commons.SqlFilter) ([]Task, error)
}
