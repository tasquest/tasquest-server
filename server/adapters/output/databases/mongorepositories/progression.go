package mongorepositories

import "tasquest.com/server/application/gamification/leveling"

type ProgressionRepository struct {
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
