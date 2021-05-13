package gorm

import (
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"tasquest.com/server/application/gamification/tasks"
	"tasquest.com/server/commons"
)

var IsTasksGormRepositoryInstanced sync.Once
var tasksGormRepositoryInstance *TasksGormRepository

type TasksGormRepository struct {
	db *gorm.DB
}

func NewTasksGormRepository(db *gorm.DB) *TasksGormRepository {
	IsTasksGormRepositoryInstanced.Do(func() {
		tasksGormRepositoryInstance = &TasksGormRepository{db: db}
		tasksGormRepositoryInstance.migrate()
	})

	return tasksGormRepositoryInstance
}

func (repo TasksGormRepository) migrate() {
	err := tasksGormRepositoryInstance.db.AutoMigrate(tasks.Task{})

	if err != nil {
		commons.IrrecoverableFailure("Failed to run SQL Migrations for the Tasks Entity", err)
	}
}

func (repo TasksGormRepository) FindByID(id uuid.UUID) (tasks.Task, error) {
	return repo.FindOneByFilter(commons.SqlFilter{"id = ?": id})
}

func (repo TasksGormRepository) FindOneByFilter(filter commons.SqlFilter) (tasks.Task, error) {
	var task tasks.Task

	tx := repo.db.First(&task, filter.ToFormattedQuery())

	if tx.Error != nil {
		return tasks.Task{}, tx.Error
	}

	return task, nil
}

func (repo TasksGormRepository) FindAllByFilter(filter commons.SqlFilter) ([]tasks.Task, error) {
	var task []tasks.Task

	tx := repo.db.Find(&task, filter.ToFormattedQuery())

	if tx.Error != nil {
		return []tasks.Task{}, tx.Error
	}

	return task, nil
}

func (repo TasksGormRepository) Save(task tasks.Task) (tasks.Task, error) {
	tx := repo.db.Create(&task)

	if tx.Error != nil {
		return tasks.Task{}, tx.Error
	}

	return task, nil
}

func (repo TasksGormRepository) Update(task tasks.Task) (tasks.Task, error) {
	tx := repo.db.Updates(&task)

	if tx.Error != nil {
		return tasks.Task{}, tx.Error
	}

	return task, nil
}

func (repo TasksGormRepository) Delete(task tasks.Task) error {
	return repo.DeleteByID(task.ID)
}

func (repo TasksGormRepository) DeleteByID(id uuid.UUID) error {
	if tx := repo.db.Delete(tasks.Task{}, id); tx.Error != nil {
		return tx.Error
	}

	return nil
}
