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
	var err error

	validate := func(apply func(command CreateLevel) error) {
		if err != nil {
			return
		}
		err = apply(command)
	}

	validate(pm.validateExperience)
	validate(pm.checkIfLevelExists)
	validate(pm.checkIfExperienceOverlaps)

	if err != nil {
		return ExpLevel{}, err
	}

	return pm.progressionPersistence.Save(ExpLevel{
		ID:    uuid.New(),
		Level: command.Level,
		From:  command.FromExp,
		To:    command.ToExp,
	})
}

func (pm ProgressionManagement) DeleteLevel(command DeleteLevel) (ExpLevel, error) {
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

func (pm ProgressionManagement) checkCharacterProgression(adventurerId uuid.UUID, character adventurers.Character) (adventurers.Character, error) {
	experience, err := pm.progressionFinder.FindLevelByExperience(character.Experience)

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

func (pm ProgressionManagement) checkIfExperienceOverlaps(command CreateLevel) error {
	if existingLevel, err := pm.progressionFinder.FindLevelByExperience(command.FromExp); err != nil {
		return errors.WithStack(err)
	} else if !cmp.Equal(existingLevel, ExpLevel{}) {
		return errors.WithStack(ErrExperienceOverlaps)
	}
	return nil
}

func (pm ProgressionManagement) checkIfLevelExists(command CreateLevel) error {
	if existingLevel, err := pm.progressionFinder.FindLevelInformation(command.Level); err != nil {
		return errors.WithStack(err)
	} else if !cmp.Equal(existingLevel, ExpLevel{}) {
		return errors.WithStack(ErrLevelAlreadyExists)
	}
	return nil
}

func (pm ProgressionManagement) validateExperience(command CreateLevel) error {
	if command.FromExp > command.ToExp {
		return errors.WithStack(ErrFromExpCannotBeHigherThanToExp)
	}
	return nil
}
