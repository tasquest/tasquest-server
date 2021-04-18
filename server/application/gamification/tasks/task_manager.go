package tasks

import (
	"emperror.dev/errors"
	"github.com/google/uuid"
	"sync"
	"tasquest.com/server/application/gamification/adventurers"
)

type TaskManager struct {
	adventurerManager adventurers.AdventurerService
	taskPersister     TaskPersistence
	taskFinder        TaskFinder
}

var TaskManagerInstanced sync.Once
var taskManagerInstance *TaskManager

func NewTaskManager(finder TaskFinder, saver TaskPersistence, management adventurers.AdventurerService) *TaskManager {
	TaskManagerInstanced.Do(func() {
		taskManagerInstance = &TaskManager{
			taskFinder:        finder,
			taskPersister:     saver,
			adventurerManager: management,
		}
	})
	return taskManagerInstance
}

func (tm *TaskManager) CreateTask(command CreateTaskCommand) (Task, error) {
	task := Task{
		Title:       command.Title,
		Description: command.Description,
		Experience:  command.Experience,
	}

	return tm.taskPersister.Save(task)
}

func (tm *TaskManager) UpdateTask(taskID uuid.UUID, command UpdateTaskCommand) (Task, error) {
	task, err := tm.taskFinder.FindByID(taskID)

	if err != nil {
		return Task{}, errors.WithStack(err)
	}

	task.Title = command.Title
	task.Description = command.Description
	task.Experience = command.Experience

	return tm.taskPersister.Update(task)
}

func (tm *TaskManager) DeleteTask(taskID uuid.UUID) (Task, error) {
	task, err := tm.taskFinder.FindByID(taskID)

	if err != nil {
		return Task{}, err
	}

	err = tm.taskPersister.Delete(task)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (tm *TaskManager) AdventurerCompletesTask(command AdventurerCompletedTaskCommand) error {
	task, err := tm.taskFinder.FindByID(command.TaskID)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = tm.adventurerManager.UpdateExperience(adventurers.UpdateExperience{
		AdventurerID: command.AdventurerID,
		Experience:   task.Experience,
	})

	return err
}
