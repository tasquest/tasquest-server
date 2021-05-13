package gorm

import (
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/commons"
)

var IsAdventurersGormRepositoryInstanced sync.Once
var adventurersGormRepositoryInstance *AdventurersGormRepository

type AdventurersGormRepository struct {
	db *gorm.DB
}

func NewAdventurersGormRepository(db *gorm.DB) *AdventurersGormRepository {
	IsAdventurersGormRepositoryInstanced.Do(func() {
		adventurersGormRepositoryInstance = &AdventurersGormRepository{db: db}
		adventurersGormRepositoryInstance.migrate()
	})

	return adventurersGormRepositoryInstance
}

func (repo AdventurersGormRepository) migrate() {
	err := adventurersGormRepositoryInstance.db.AutoMigrate(adventurers.Adventurer{})

	if err != nil {
		commons.IrrecoverableFailure("Failed to run SQL Migrations for the Adventurer Entity", err)
	}
}

func (repo AdventurersGormRepository) FindByID(id uuid.UUID) (adventurers.Adventurer, error) {
	return repo.FindOneByFilter(commons.SqlFilter{"id = ?": id})
}

func (repo AdventurersGormRepository) FindByUser(userID uuid.UUID) (adventurers.Adventurer, error) {
	return repo.FindOneByFilter(commons.SqlFilter{"user_id = ?": userID})
}

func (repo AdventurersGormRepository) FindOneByFilter(filter commons.SqlFilter) (adventurers.Adventurer, error) {
	var adventurer adventurers.Adventurer

	tx := repo.db.First(&adventurer, filter.ToFormattedQuery())

	if tx.Error != nil {
		return adventurers.Adventurer{}, tx.Error
	}

	return adventurer, nil
}

func (repo AdventurersGormRepository) FindAllByFilter(filter commons.SqlFilter) ([]adventurers.Adventurer, error) {
	var adventurer []adventurers.Adventurer

	tx := repo.db.Find(&adventurer, filter.ToFormattedQuery())

	if tx.Error != nil {
		return []adventurers.Adventurer{}, tx.Error
	}

	return adventurer, nil
}

func (repo AdventurersGormRepository) Save(adventurer adventurers.Adventurer) (adventurers.Adventurer, error) {
	tx := repo.db.Create(&adventurer)

	if tx.Error != nil {
		return adventurers.Adventurer{}, tx.Error
	}

	return adventurer, nil
}

func (repo AdventurersGormRepository) Update(adventurer adventurers.Adventurer) (adventurers.Adventurer, error) {
	tx := repo.db.Updates(&adventurer)

	if tx.Error != nil {
		return adventurers.Adventurer{}, tx.Error
	}

	return adventurer, nil
}
