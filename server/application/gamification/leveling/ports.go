package leveling

import (
	"github.com/google/uuid"

	"tasquest.com/server/commons"
)

type ProgressionService interface {
	CreateLevel(command CreateLevel) (ExpLevel, error)
	DeleteHighestLevel() (ExpLevel, error)
	AwardExperience(command AwardExperience) error
}

type ProgressionPersistence interface {
	Save(level ExpLevel) (ExpLevel, error)
	Delete(level ExpLevel) (ExpLevel, error)
}

type ProgressionFinder interface {
	FindLevelInformation(level int) (ExpLevel, error)
	FindLatestLevel() (ExpLevel, error)
	FindLevelForExperience(experience int64) (ExpLevel, error)
	FindByID(id uuid.UUID) (ExpLevel, error)
	FindOneByFilter(filter commons.SqlFilter) (ExpLevel, error)
}
