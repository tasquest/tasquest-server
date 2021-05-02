package mongorepositories

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"

	"tasquest.com/server/application/gamification/leveling"
)

var IsProgressionRepositoryInstanced = sync.Once{}
var progressionRepositoryInstance *ProgressionRepository

type ProgressionRepository struct {
	collection *mongo.Collection
}

func NewProgressionRepository(dbClient *mongo.Database) *ProgressionRepository {
	IsProgressionRepositoryInstanced.Do(func() {
		progressionRepositoryInstance = &ProgressionRepository{
			collection: dbClient.Collection("progressions"),
		}
	})
	return progressionRepositoryInstance
}

func (p ProgressionRepository) Save(level leveling.ExpLevel) (leveling.ExpLevel, error) {
	panic("implement me")
}

func (p ProgressionRepository) Update(level leveling.ExpLevel) (leveling.ExpLevel, error) {
	panic("implement me")
}

func (p ProgressionRepository) Delete(level leveling.ExpLevel) (leveling.ExpLevel, error) {
	panic("implement me")
}

func (p ProgressionRepository) FindLevelInformation(level int) (leveling.ExpLevel, error) {
	panic("implement me")
}

func (p ProgressionRepository) FindLevelByExperience(experience int64) (leveling.ExpLevel, error) {
	panic("implement me")
}
