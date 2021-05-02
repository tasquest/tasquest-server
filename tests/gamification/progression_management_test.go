package gamification

import (
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/application/gamification/leveling"
	"tasquest.com/tests/mocks"
)

type ProgressionManagementTestSuite struct {
	suite.Suite
	progressionFinder      *mocks.ProgressionFinder
	progressionPersistence *mocks.ProgressionPersistence
	adventurerService      *mocks.AdventurerService
	adventurerFinder       *mocks.AdventurerFinder
	eventPublisher         *mocks.Publisher
}

func (suite *ProgressionManagementTestSuite) SetupTest() {
	leveling.IsProgressionManagementInstanced = sync.Once{}
	suite.progressionFinder = new(mocks.ProgressionFinder)
	suite.progressionPersistence = new(mocks.ProgressionPersistence)
	suite.adventurerService = new(mocks.AdventurerService)
	suite.adventurerFinder = new(mocks.AdventurerFinder)
	suite.eventPublisher = new(mocks.Publisher)
}

func (suite *ProgressionManagementTestSuite) checkMockAssertions() {
	suite.progressionFinder.AssertExpectations(suite.T())
	suite.progressionPersistence.AssertExpectations(suite.T())
	suite.adventurerFinder.AssertExpectations(suite.T())
	suite.adventurerService.AssertExpectations(suite.T())
	suite.eventPublisher.AssertExpectations(suite.T())
}

func (suite *ProgressionManagementTestSuite) TestLevelCreatedSuccessfully() {
	progressionManagement := leveling.NewProgressionManagement(
		suite.progressionFinder,
		suite.progressionPersistence,
		suite.adventurerService,
		suite.adventurerFinder,
		suite.eventPublisher,
	)

	command := leveling.CreateLevel{
		Level:   1,
		FromExp: 10,
		ToExp:   20,
	}

	expectedResponse := leveling.ExpLevel{
		ID:    uuid.New(),
		Level: 1,
		From:  10,
		To:    20,
	}

	suite.progressionFinder.On("FindLevelInformation", 1).Return(leveling.ExpLevel{}, nil)
	suite.progressionFinder.On("FindLevelByExperience", int64(10)).Return(leveling.ExpLevel{}, nil)
	suite.progressionPersistence.On("Save", mock.Anything).Return(expectedResponse, nil)

	expLevel, err := progressionManagement.CreateLevel(command)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedResponse, expLevel)

	suite.checkMockAssertions()
}

func (suite *ProgressionManagementTestSuite) TestLevelFromExperienceLargerThanToExperience() {
	progressionManagement := leveling.NewProgressionManagement(
		suite.progressionFinder,
		suite.progressionPersistence,
		suite.adventurerService,
		suite.adventurerFinder,
		suite.eventPublisher,
	)

	command := leveling.CreateLevel{
		Level:   1,
		FromExp: 20,
		ToExp:   10,
	}

	_, err := progressionManagement.CreateLevel(command)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), leveling.ErrFromExpCannotBeHigherThanToExp.Message, err.Error())

	suite.checkMockAssertions()
}

func (suite *ProgressionManagementTestSuite) TestAwardedExperienceLevelsCharacterUp() {
	progressionManagement := leveling.NewProgressionManagement(
		suite.progressionFinder,
		suite.progressionPersistence,
		suite.adventurerService,
		suite.adventurerFinder,
		suite.eventPublisher,
	)

	adventurerId := uuid.New()

	existingCharacter := adventurers.Character{
		CharName:   "Char",
		Title:      "The Unbroken",
		Level:      1,
		Experience: 10,
		Equipment:  adventurers.CharacterEquipment{},
		Guilds:     nil,
	}

	expectedCharacter := existingCharacter.IncreaseExperience(10).NewLevel(2)

	adventurer := adventurers.Adventurer{
		ID:        adventurerId,
		UserID:    uuid.New(),
		Name:      "Test",
		Surname:   "Test",
		Birthday:  time.Time{},
		Character: existingCharacter,
	}

	command := leveling.AwardExperience{
		AdventurerID: adventurerId,
		Experience:   10,
	}

	suite.adventurerFinder.On("FindByID", adventurerId).Return(adventurer, nil)

	suite.progressionFinder.On("FindLevelByExperience", int64(20)).Return(leveling.ExpLevel{
		ID:    uuid.UUID{},
		Level: 2,
		From:  15,
		To:    30,
	}, nil)

	suite.eventPublisher.On("Publish", leveling.AdventurerLevelingTopic, leveling.AdventurerLevelingEvent{
		AdventurerID: adventurerId,
		OldLevel:     1,
		NewLevel:     2,
	}).Return(leveling.AdventurerLevelingEvent{}, nil)

	suite.adventurerService.On("UpdateCharacter", adventurerId, expectedCharacter).Return(adventurers.Adventurer{}, nil)

	err := progressionManagement.AwardExperience(command)

	assert.Nil(suite.T(), err)

	suite.checkMockAssertions()
}

func (suite *ProgressionManagementTestSuite) TestAwardedExperienceDoesNotLevelsCharacterUp() {
	progressionManagement := leveling.NewProgressionManagement(
		suite.progressionFinder,
		suite.progressionPersistence,
		suite.adventurerService,
		suite.adventurerFinder,
		suite.eventPublisher,
	)

	adventurerId := uuid.New()

	existingCharacter := adventurers.Character{
		CharName:   "Char",
		Title:      "The Unbroken",
		Level:      1,
		Experience: 10,
		Equipment:  adventurers.CharacterEquipment{},
		Guilds:     nil,
	}

	expectedCharacter := existingCharacter.IncreaseExperience(1)

	adventurer := adventurers.Adventurer{
		ID:        adventurerId,
		UserID:    uuid.New(),
		Name:      "Test",
		Surname:   "Test",
		Birthday:  time.Time{},
		Character: existingCharacter,
	}

	command := leveling.AwardExperience{
		AdventurerID: adventurerId,
		Experience:   1,
	}

	suite.adventurerFinder.On("FindByID", adventurerId).Return(adventurer, nil)

	suite.progressionFinder.On("FindLevelByExperience", int64(11)).Return(leveling.ExpLevel{
		ID:    uuid.UUID{},
		Level: 1,
		From:  10,
		To:    20,
	}, nil)

	suite.adventurerService.On("UpdateCharacter", adventurerId, expectedCharacter).Return(adventurers.Adventurer{}, nil)

	err := progressionManagement.AwardExperience(command)

	assert.Nil(suite.T(), err)

	suite.checkMockAssertions()
}

func TestRunProgressionManagementSuite(t *testing.T) {
	suite.Run(t, new(ProgressionManagementTestSuite))
}
