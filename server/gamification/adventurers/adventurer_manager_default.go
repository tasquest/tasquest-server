package adventurers

import (
	"emperror.dev/errors"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"tasquest.com/server/security"
)

var IsAdventurerManagerInstanced sync.Once
var adventurerManagerInstance *DefaultAdventurerManager

type DefaultAdventurerManager struct {
	adventurerRepository AdventurerRepository
	userManagement       security.UserManagement
}

func ProvideDefaultAdventurerManager(repository AdventurerRepository, management security.UserManagement) *DefaultAdventurerManager {
	IsAdventurerManagerInstanced.Do(func() {
		adventurerManagerInstance = &DefaultAdventurerManager{
			adventurerRepository: repository,
			userManagement:       management,
		}
	})
	return adventurerManagerInstance
}

func (pm *DefaultAdventurerManager) CreateAdventurer(command CreateAdventurer) (Adventurer, error) {
	_, profErr := pm.FetchAdventurerForUser(command.UserID)

	if profErr != nil {
		return Adventurer{}, errors.WithStack(profErr)
	}

	userID, _ := primitive.ObjectIDFromHex(command.UserID)

	profile := Adventurer{
		UserID:   userID,
		Name:     command.Name,
		Surname:  command.Surname,
		Birthday: primitive.NewDateTimeFromTime(command.Birthday),
	}

	return pm.adventurerRepository.Save(profile)
}

func (pm *DefaultAdventurerManager) FetchAdventurerForUser(userID string) (Adventurer, error) {
	_, err := pm.userManagement.FetchUser(userID)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	existingProfile, err := pm.adventurerRepository.FindByUser(userID)

	if !cmp.Equal(existingProfile, Adventurer{}) {
		return Adventurer{}, errors.WithStack(ErrAdventurerAlreadyExists)
	}

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	return existingProfile, nil
}

func (pm *DefaultAdventurerManager) UpdateAdventurer(adventurerID string, command UpdateAdventurer) (Adventurer, error) {
	adventurer, err := pm.adventurerRepository.FindByID(adventurerID)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	adventurer.Birthday = primitive.NewDateTimeFromTime(command.Birthday)
	adventurer.Surname = command.Surname
	adventurer.Name = command.Name

	updated, err := pm.adventurerRepository.Update(adventurer)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	return updated, nil
}

func (pm *DefaultAdventurerManager) UpdateAdventurerExperience(command UpdateExperience) (Adventurer, error) {
	adventurer, err := pm.adventurerRepository.FindByID(command.AdventurerID)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	adventurer.Character.Experience += command.Experience

	updatedAdventurer, err := pm.adventurerRepository.Update(adventurer)

	if err != nil {
		return Adventurer{}, errors.WithStack(err)
	}

	return updatedAdventurer, nil
}
