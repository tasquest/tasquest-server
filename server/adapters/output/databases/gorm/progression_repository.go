package gorm

import (
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"tasquest.com/server/application/gamification/leveling"
	"tasquest.com/server/commons"
)

var IsProgressionGormRepositoryInstanced sync.Once
var progressionGormRepositoryInstance *ProgressionGormRepository

type ProgressionGormRepository struct {
	db *gorm.DB
}

func NewProgressionGormRepository(db *gorm.DB) *ProgressionGormRepository {
	IsProgressionGormRepositoryInstanced.Do(func() {
		progressionGormRepositoryInstance = &ProgressionGormRepository{db: db}
		progressionGormRepositoryInstance.migrate()
	})
	return progressionGormRepositoryInstance
}

func (repo ProgressionGormRepository) migrate() {
	err := progressionGormRepositoryInstance.db.AutoMigrate(leveling.ExpLevel{})

	if err != nil {
		commons.IrrecoverableFailure("Failed to run SQL Migrations for the ExpLevel Entity", err)
	}
}

func (repo ProgressionGormRepository) FindLevelInformation(level int) (leveling.ExpLevel, error) {
	return repo.FindOneByFilter(commons.SqlFilter{"level = ?": level})
}

func (repo ProgressionGormRepository) FindLatestLevel() (leveling.ExpLevel, error) {
	var expLevel leveling.ExpLevel

	if tx := repo.db.Raw("SELECT * FROM exp_levels WHERE level = (SELECT MAX(level) FROM exp_levels)").Scan(&expLevel); tx.Error != nil {
		return leveling.ExpLevel{}, tx.Error
	}

	return expLevel, nil
}

func (repo ProgressionGormRepository) FindLevelForExperience(experience int64) (leveling.ExpLevel, error) {
	return repo.FindOneByFilter(commons.SqlFilter{"start_exp < ?": experience, "AND end_exp > ?": experience})
}

func (repo ProgressionGormRepository) FindByID(id uuid.UUID) (leveling.ExpLevel, error) {
	return repo.FindOneByFilter(commons.SqlFilter{"id = ?": id})
}

func (repo ProgressionGormRepository) FindOneByFilter(filter commons.SqlFilter) (leveling.ExpLevel, error) {
	var expLevel leveling.ExpLevel
	params := filter.ToFormattedQuery()

	if tx := repo.db.Where(params.Query, params.Params...).First(&expLevel); tx.Error != nil {
		return leveling.ExpLevel{}, tx.Error
	}

	return expLevel, nil
}

func (repo ProgressionGormRepository) Save(level leveling.ExpLevel) (leveling.ExpLevel, error) {
	if tx := repo.db.Create(&level); tx.Error != nil {
		return leveling.ExpLevel{}, tx.Error
	}
	return level, nil
}

func (repo ProgressionGormRepository) Delete(level leveling.ExpLevel) (leveling.ExpLevel, error) {
	if tx := repo.db.Delete(leveling.ExpLevel{}, level.ID); tx.Error != nil {
		return leveling.ExpLevel{}, tx.Error
	}

	return level, nil
}
