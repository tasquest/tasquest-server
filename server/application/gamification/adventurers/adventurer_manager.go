package adventurers

import (
	"sync"

	"emperror.dev/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"

	"tasquest.com/server/application/security"
)

var IsAdventurerManagerInstanced sync.Once
var adventurerManagerInstance *AdventurerManager

type AdventurerManager struct {
	adventurerFinder AdventurerFinder
	adventurerSaver  AdventurerPersistence
	userManagement   security.UserService
}

func NewAdventurerManager(finder AdventurerFinder, saver AdventurerPersistence, management security.UserService) *AdventurerManager {
	IsAdventurerManagerInstanced.Do(func() {
		adventurerManagerInstance = &AdventurerManager{
			adventurerFinder: finder,
			adventurerSaver:  saver,
			userManagement:   management,
		}
	})
	return adventurerManagerInstance
}

func (am *AdventurerManager) CreateAdventurer(command CreateAdventurer) (Adventurer, error) {
	existingAdventurer, err := am.adventurerFinder.FindByUser(command.UserID)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	if !cmp.Equal(existingAdventurer, Adventurer{}) {
		return Adventurer{}, errors.WithStack(ErrAdventurerAlreadyExists)
	}

	adventurer := Adventurer{
		UserID:   command.UserID,
		Name:     command.Name,
		Surname:  command.Surname,
		Birthday: command.Birthday,
	}

	return am.adventurerSaver.Save(adventurer)
}

func (am *AdventurerManager) UpdateAdventurer(adventurerID uuid.UUID, command UpdateAdventurer) (Adventurer, error) {
	adventurer, err := am.adventurerFinder.FindByID(adventurerID)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	adventurer.Birthday = command.Birthday
	adventurer.Surname = command.Surname
	adventurer.Name = command.Name

	return am.adventurerSaver.Update(adventurer)
}

func (am *AdventurerManager) UpdateCharacter(adventurerID uuid.UUID, character Character) (Adventurer, error) {
	panic("implement me")
}
