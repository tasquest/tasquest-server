package leveling

type ProgressionService interface {
	CreateLevel(command CreateLevel) (ExpLevel, error)
	DeleteLevel(command DeleteLevel) (ExpLevel, error)
	AwardExperience(command AwardExperience) error
}

type ProgressionPersistence interface {
	Save(level ExpLevel) (ExpLevel, error)
	Update(level ExpLevel) (ExpLevel, error)
	Delete(level ExpLevel) (ExpLevel, error)
}

type ProgressionFinder interface {
	FindLevelInformation(level int) (ExpLevel, error)
	FindLevelByExperience(experience int64) (ExpLevel, error)
}
