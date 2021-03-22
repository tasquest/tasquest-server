package tasks

import (
	"emperror.dev/errors"
	"sync"
	"tasquest.com/server/gamification/adventurers"
)

type TaskManagerImpl struct {
	adventurerManager adventurers.AdventurerManager
	taskRepository    TaskRepository
}

var TaskManagerOnce sync.Once
var taskManagerInstance *TaskManagerImpl

func ProvideDefaultTaskManager(taskRepository TaskRepository, profileManager adventurers.AdventurerManager) *TaskManagerImpl {
	TaskManagerOnce.Do(func() {
		taskManagerInstance = &TaskManagerImpl{
			adventurerManager: profileManager,
			taskRepository:    taskRepository,
		}
	})
	return taskManagerInstance
}

func (t *TaskManagerImpl) CreateTask(command CreateTaskCommand) (Task, error) {
	task := Task{
		Title:       command.Title,
		Description: command.Description,
		Experience:  command.Experience,
	}

	return t.taskRepository.Save(task)
}

func (t *TaskManagerImpl) UpdateTask(taskID string, command UpdateTaskCommand) (Task, error) {
	task, err := t.taskRepository.FindByID(taskID)

	if err != nil {
		return Task{}, errors.WithStack(err)
	}

	task.Title = command.Title
	task.Description = command.Description
	task.Experience = command.Experience

	return t.taskRepository.Update(task)
}

func (t *TaskManagerImpl) DeleteTask(taskID string) (Task, error) {
	task, err := t.taskRepository.FindByID(taskID)

	if err != nil {
		return Task{}, err
	}

	err = t.taskRepository.Delete(task)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (t *TaskManagerImpl) AdventurerCompletesTask(command AdventurerCompletedTaskCommand) error {
	task, err := t.taskRepository.FindByID(command.TaskID)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = t.adventurerManager.UpdateAdventurerExperience(adventurers.UpdateExperience{
		AdventurerID: command.AdventurerID,
		Experience:   task.Experience,
	})

	return err
}
