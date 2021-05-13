package leveling

import (
	"sync"

	"emperror.dev/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"

	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/infra/events"
)

var IsProgressionManagementInstanced sync.Once
var progressionManagementInstance *ProgressionManagement

type ProgressionManagement struct {
	progressionFinder      ProgressionFinder
	progressionPersistence ProgressionPersistence
	adventurerService      adventurers.AdventurerService
	adventurerFinder       adventurers.AdventurerFinder
	publisher              events.Publisher
}

func NewProgressionManagement(
	progressionFinder ProgressionFinder,
	progressionPersistence ProgressionPersistence,
	adventurerService adventurers.AdventurerService,
	adventurerFinder adventurers.AdventurerFinder,
	publisher events.Publisher,
) *ProgressionManagement {
	IsProgressionManagementInstanced.Do(func() {
		progressionManagementInstance = &ProgressionManagement{
			progressionFinder:      progressionFinder,
			progressionPersistence: progressionPersistence,
			adventurerService:      adventurerService,
			adventurerFinder:       adventurerFinder,
			publisher:              publisher,
		}
	})
	return progressionManagementInstance
}

func (pm ProgressionManagement) CreateLevel(command CreateLevel) (ExpLevel, error) {
	newLevel, latestExperience, err := pm.calculateNextLevel(command)

	if err != nil {
		return ExpLevel{}, err
	}

	return pm.progressionPersistence.Save(ExpLevel{
		ID:       uuid.New(),
		Level:    newLevel,
		StartExp: latestExperience,
		EndExp:   command.NewTopExp,
	})
}

func (pm ProgressionManagement) DeleteHighestLevel() (ExpLevel, error) {
	panic("implement me")
}

func (pm ProgressionManagement) AwardExperience(command AwardExperience) error {
	adventurer, err := pm.adventurerFinder.FindByID(command.AdventurerID)

	if err != nil {
		return errors.WithStack(err)
	}

	character := adventurer.Character.IncreaseExperience(command.Experience)
	character, err = pm.checkCharacterProgression(command.AdventurerID, character)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = pm.adventurerService.UpdateCharacter(command.AdventurerID, character)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Helper Functions

func (pm ProgressionManagement) checkCharacterProgression(
	adventurerId uuid.UUID,
	character adventurers.Character,
) (adventurers.Character, error) {
	experience, err := pm.progressionFinder.FindLevelForExperience(character.Experience)

	if err != nil {
		return adventurers.Character{}, errors.WithStack(err)
	}

	if experience.Level != character.Level {
		_, err := pm.publisher.Publish(AdventurerLevelingTopic, AdventurerLevelingEvent{
			AdventurerID: adventurerId,
			OldLevel:     character.Level,
			NewLevel:     experience.Level,
		})

		if err != nil {
			return adventurers.Character{}, errors.WithStack(err)
		}

		return character.NewLevel(experience.Level), nil
	}

	return character, nil
}

func (pm ProgressionManagement) calculateNextLevel(command CreateLevel) (int, int64, error) {
	levelBefore, err := pm.progressionFinder.FindLatestLevel()

	if err != nil {
		return -1, -1, err
	}

	if levelBefore.EndExp > command.NewTopExp {
		return -1, -1, errors.WithStack(ErrNewLevelCantBeLowerThanLargestExistingLevel)
	}

	if !cmp.Equal(levelBefore, ExpLevel{}) {
		return levelBefore.Level + 1, levelBefore.EndExp + 1, nil
	}

	return 1, 1, nil
}
