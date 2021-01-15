package profiles

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tasquest.com/server/gamification/adventurers"
	"tasquest.com/server/security"
	"tasquest.com/tests/mocks"
	"testing"
	"time"
)

func TestCreateProfileSuccessfully(t *testing.T) {
	profileRepository := new(mocks.ProfileRepository)
	userManagement := new(mocks.UserManagement)
	profileManager := adventurers.DefaultAdventurerManager{
		adventurerRepository: profileRepository,
		userManagement:       userManagement,
	}

	command := adventurers.CreateAdventurer{
		UserID:   "1",
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	userID, _ := primitive.ObjectIDFromHex("1")

	profile := adventurers.Adventurer{
		ID:       primitive.ObjectID{},
		UserID:   userID,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: primitive.NewDateTimeFromTime(command.Birthday),
	}

	userManagement.On("FetchUser", "1").Return(security.User{}, nil)
	profileRepository.On("FindByUser", "1").Return(adventurers.Adventurer{}, nil)
	profileRepository.On("Save", mock.Anything).Return(profile, nil)

	createdProfile, err := profileManager.CreateProfile(command)

	assert.Nil(t, err)
	assert.NotNil(t, createdProfile)
	assert.Equal(t, profile, createdProfile)

	profileRepository.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateProfileNoUser(t *testing.T) {
	profileRepository := new(mocks.ProfileRepository)
	userManagement := new(mocks.UserManagement)
	profileManager := adventurers.DefaultAdventurerManager{
		adventurerRepository: profileRepository,
		userManagement:       userManagement,
	}

	command := adventurers.CreateAdventurer{
		UserID:   "1",
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	userManagement.On("FetchUser", "1").Return(security.User{}, security.ErrUserNotFound)

	createdProfile, err := profileManager.CreateProfile(command)

	assert.NotNil(t, err)
	assert.Equal(t, security.ErrUserNotFound.Message, err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdProfile)

	profileRepository.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateProfileAlreadyExists(t *testing.T) {
	profileRepository := new(mocks.ProfileRepository)
	userManagement := new(mocks.UserManagement)
	profileManager := adventurers.DefaultAdventurerManager{
		adventurerRepository: profileRepository,
		userManagement:       userManagement,
	}

	command := adventurers.CreateAdventurer{
		UserID:   "1",
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	userID, _ := primitive.ObjectIDFromHex("1")

	existingProfile := adventurers.Adventurer{
		ID:       primitive.ObjectID{},
		UserID:   userID,
		Name:     "Test",
		Surname:  "Tester",
		Birthday: 0,
	}

	userManagement.On("FetchUser", "1").Return(security.User{}, nil)
	profileRepository.On("FindByUser", "1").Return(existingProfile, nil)

	createdProfile, err := profileManager.CreateProfile(command)

	assert.NotNil(t, err)
	assert.Equal(t, adventurers.ErrAdventurerAlreadyExists.Message, err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdProfile)

	profileRepository.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}

func TestCreateProfileFailedToFetchProfile(t *testing.T) {
	profileRepository := new(mocks.ProfileRepository)
	userManagement := new(mocks.UserManagement)
	profileManager := adventurers.DefaultAdventurerManager{
		adventurerRepository: profileRepository,
		userManagement:       userManagement,
	}

	command := adventurers.CreateAdventurer{
		UserID:   "1",
		Name:     "Test",
		Surname:  "Tester",
		Birthday: time.Now(),
	}

	userManagement.On("FetchUser", "1").Return(security.User{}, nil)
	profileRepository.On("FindByUser", "1").Return(adventurers.Adventurer{}, adventurers.ErrFailedToFetchAdventurer)

	createdProfile, err := profileManager.CreateProfile(command)

	assert.NotNil(t, err)
	assert.Equal(t, "Failed to fetch profile: An unexpected error occurred", err.Error())
	assert.Equal(t, adventurers.Adventurer{}, createdProfile)

	profileRepository.AssertExpectations(t)
	userManagement.AssertExpectations(t)
}
