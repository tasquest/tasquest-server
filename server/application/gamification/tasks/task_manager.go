package tasks

import (
	"sync"

	"emperror.dev/errors"
	"github.com/google/uuid"

	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/infra/events"
)

type TaskManager struct {
	adventurerFinder adventurers.AdventurerFinder
	taskPersistence  TaskPersistence
	taskFinder       TaskFinder
	eventPublisher   events.Publisher
}

var TaskManagerInstanced sync.Once
var taskManagerInstance *TaskManager

func NewTaskManager(finder TaskFinder, saver TaskPersistence, management adventurers.AdventurerFinder) *TaskManager {
	TaskManagerInstanced.Do(func() {
		taskManagerInstance = &TaskManager{
			taskFinder:       finder,
			taskPersistence:  saver,
			adventurerFinder: management,
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

	return tm.taskPersistence.Save(task)
}

func (tm *TaskManager) UpdateTask(taskID uuid.UUID, command UpdateTaskCommand) (Task, error) {
	task, err := tm.taskFinder.FindByID(taskID)

	if err != nil {
		return Task{}, errors.WithStack(err)
	}

	task.Title = command.Title
	task.Description = command.Description
	task.Experience = command.Experience

	return tm.taskPersistence.Update(task)
}

func (tm *TaskManager) DeleteTask(taskID uuid.UUID) (Task, error) {
	task, err := tm.taskFinder.FindByID(taskID)

	if err != nil {
		return Task{}, err
	}

	err = tm.taskPersistence.Delete(task)

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

	_, err = tm.adventurerFinder.FindByID(command.AdventurerID)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = tm.eventPublisher.Publish(AdventurerTaskTopic, AdventurerCompletedTaskEvent{
		AdventurerID:      command.AdventurerID,
		TaskId:            command.TaskID,
		ExperienceAwarded: task.Experience,
	})

	return errors.WithStack(err)
}
