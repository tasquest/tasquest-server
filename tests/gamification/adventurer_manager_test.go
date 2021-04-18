package gamification

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/application/security"
	"tasquest.com/tests/mocks"
	"testing"
	"time"
)

func TestCreateProfileSuccessfully(t *testing.T) {
	adventurers.IsAdventurerManagerInstanced = sync.Once{}
	adventurerFinder := new(mocks.AdventurerFinder)
	adventurerPersistence := new(mocks.AdventurerPersistence)
	userManagement := new(mocks.UserService)
	adventurerManager := adventurers.NewAdventurerManager(adventurerFinder, adventurerPersistence, userManagement)
	userId, _ := uuid.NewUUID()
	adventurerId, _ := uuid.NewUUID()

	command := adventurers.CreateAdventurer{
		UserID:   userId,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	adventurer := adventurers.Adventurer{
		ID:       adventurerId,
		UserID:   userId,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: command.Birthday,
	}

	adventurerFinder.On("FindByUser", userId).Return(adventurers.Adventurer{}, nil)
	adventurerPersistence.On("Save", mock.Anything).Return(adventurer, nil)

	createdAdventurer, err := adventurerManager.CreateAdventurer(command)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdventurer)
	assert.Equal(t, adventurer, createdAdventurer)

	adventurerFinder.AssertExpectations(t)
	adventurerPersistence.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateAdventurerNoUser(t *testing.T) {
	adventurers.IsAdventurerManagerInstanced = sync.Once{}
	adventurerFinder := new(mocks.AdventurerFinder)
	adventurerPersistence := new(mocks.AdventurerPersistence)
	userManagement := new(mocks.UserService)
	adventurerManager := adventurers.NewAdventurerManager(adventurerFinder, adventurerPersistence, userManagement)
	userId, _ := uuid.NewUUID()

	command := adventurers.CreateAdventurer{
		UserID:   userId,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	adventurerFinder.On("FindByUser", userId).Return(adventurers.Adventurer{}, security.ErrUserNotFound)

	createdAdventurer, err := adventurerManager.CreateAdventurer(command)

	assert.NotNil(t, err)
	assert.Equal(t, security.ErrUserNotFound.Message, err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdAdventurer)

	adventurerFinder.AssertExpectations(t)
	adventurerPersistence.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateAdventurerAlreadyExists(t *testing.T) {
	adventurers.IsAdventurerManagerInstanced = sync.Once{}
	adventurerFinder := new(mocks.AdventurerFinder)
	adventurerPersistence := new(mocks.AdventurerPersistence)
	userManagement := new(mocks.UserService)
	adventurerManager := adventurers.NewAdventurerManager(adventurerFinder, adventurerPersistence, userManagement)
	userId, _ := uuid.NewUUID()
	adventurerId, _ := uuid.NewUUID()

	command := adventurers.CreateAdventurer{
		UserID:   userId,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	existingAdventurer := adventurers.Adventurer{
		ID:       adventurerId,
		UserID:   userId,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	adventurerFinder.On("FindByUser", userId).Return(existingAdventurer, nil)

	createdProfile, err := adventurerManager.CreateAdventurer(command)

	assert.NotNil(t, err)
	assert.Equal(t, adventurers.ErrAdventurerAlreadyExists.Message, err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdProfile)

	adventurerFinder.AssertExpectations(t)
	adventurerPersistence.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateAdventurerFailedToFetchAdventurer(t *testing.T) {
	adventurers.IsAdventurerManagerInstanced = sync.Once{}
	adventurerFinder := new(mocks.AdventurerFinder)
	adventurerPersistence := new(mocks.AdventurerPersistence)
	userManagement := new(mocks.UserService)
	adventurerManager := adventurers.NewAdventurerManager(adventurerFinder, adventurerPersistence, userManagement)
	userId, _ := uuid.NewUUID()

	command := adventurers.CreateAdventurer{
		UserID:   userId,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	adventurerFinder.On("FindByUser", userId).Return(adventurers.Adventurer{}, adventurers.ErrFailedToFetchAdventurer)

	createdAdventurer, err := adventurerManager.CreateAdventurer(command)

	assert.NotNil(t, err)
	assert.Equal(t, "An unexpected error occurred", err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdAdventurer)

	adventurerFinder.AssertExpectations(t)
	adventurerPersistence.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}
