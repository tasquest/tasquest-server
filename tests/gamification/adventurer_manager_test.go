package gamification

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"tasquest.com/server/gamification/adventurers"
	"tasquest.com/server/security"
	"tasquest.com/tests/mocks"
	"testing"
	"time"
)

func TestCreateProfileSuccessfully(t *testing.T) {
	adventurers.IsAdventurerManagerInstanced = sync.Once{}
	adventurerRepository := new(mocks.AdventurerRepository)
	userManagement := new(mocks.UserManagement)
	adventurerManager := adventurers.ProvideDefaultAdventurerManager(adventurerRepository, userManagement)

	command := adventurers.CreateAdventurer{
		UserID:   "1",
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	userID, _ := primitive.ObjectIDFromHex("1")

	adventurer := adventurers.Adventurer{
		ID:       primitive.ObjectID{},
		UserID:   userID,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: primitive.NewDateTimeFromTime(command.Birthday),
	}

	userManagement.On("FetchUser", "1").Return(security.User{}, nil)
	adventurerRepository.On("FindByUser", "1").Return(adventurers.Adventurer{}, nil)
	adventurerRepository.On("Save", mock.Anything).Return(adventurer, nil)

	createdAdventurer, err := adventurerManager.CreateAdventurer(command)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdventurer)
	assert.Equal(t, adventurer, createdAdventurer)

	adventurerRepository.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateAdventurerNoUser(t *testing.T) {
	adventurers.IsAdventurerManagerInstanced = sync.Once{}
	adventurerRepository := new(mocks.AdventurerRepository)
	userManagement := new(mocks.UserManagement)
	adventurerManager := adventurers.ProvideDefaultAdventurerManager(adventurerRepository, userManagement)

	command := adventurers.CreateAdventurer{
		UserID:   "1",
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	userManagement.On("FetchUser", "1").Return(security.User{}, security.ErrUserNotFound)

	createdAdventurer, err := adventurerManager.CreateAdventurer(command)

	assert.NotNil(t, err)
	assert.Equal(t, security.ErrUserNotFound.Message, err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdAdventurer)

	adventurerRepository.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateAdventurerAlreadyExists(t *testing.T) {
	adventurers.IsAdventurerManagerInstanced = sync.Once{}
	adventurerRepository := new(mocks.AdventurerRepository)
	userManagement := new(mocks.UserManagement)
	adventurerManager := adventurers.ProvideDefaultAdventurerManager(adventurerRepository, userManagement)

	command := adventurers.CreateAdventurer{
		UserID:   "1",
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	userID, _ := primitive.ObjectIDFromHex("1")

	existingAdventurer := adventurers.Adventurer{
		ID:       primitive.ObjectID{},
		UserID:   userID,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: 0,
	}

	userManagement.On("FetchUser", "1").Return(security.User{}, nil)
	adventurerRepository.On("FindByUser", "1").Return(existingAdventurer, nil)

	createdProfile, err := adventurerManager.CreateAdventurer(command)

	assert.NotNil(t, err)
	assert.Equal(t, adventurers.ErrAdventurerAlreadyExists.Message, err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdProfile)

	adventurerRepository.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateAdventurerFailedToFetchAdventurer(t *testing.T) {
	adventurers.IsAdventurerManagerInstanced = sync.Once{}
	adventurerRepository := new(mocks.AdventurerRepository)
	userManagement := new(mocks.UserManagement)
	adventurerManager := adventurers.ProvideDefaultAdventurerManager(adventurerRepository, userManagement)

	command := adventurers.CreateAdventurer{
		UserID:   "1",
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	userManagement.On("FetchUser", "1").Return(security.User{}, nil)
	adventurerRepository.On("FindByUser", "1").Return(adventurers.Adventurer{}, adventurers.ErrFailedToFetchAdventurer)

	createdAdventurer, err := adventurerManager.CreateAdventurer(command)

	assert.NotNil(t, err)
	assert.Equal(t, "An unexpected error occurred", err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdAdventurer)

	adventurerRepository.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}
