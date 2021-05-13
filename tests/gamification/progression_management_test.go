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
	// Dependencies
	progressionFinder      *mocks.ProgressionFinder
	progressionPersistence *mocks.ProgressionPersistence
	adventurerService      *mocks.AdventurerService
	adventurerFinder       *mocks.AdventurerFinder
	eventPublisher         *mocks.Publisher
	// InjectTo
	progressionManagement *leveling.ProgressionManagement
}

func (suite *ProgressionManagementTestSuite) SetupTest() {
	leveling.IsProgressionManagementInstanced = sync.Once{}
	suite.progressionFinder = new(mocks.ProgressionFinder)
	suite.progressionPersistence = new(mocks.ProgressionPersistence)
	suite.adventurerService = new(mocks.AdventurerService)
	suite.adventurerFinder = new(mocks.AdventurerFinder)
	suite.eventPublisher = new(mocks.Publisher)
	suite.progressionManagement = leveling.NewProgressionManagement(
		suite.progressionFinder,
		suite.progressionPersistence,
		suite.adventurerService,
		suite.adventurerFinder,
		suite.eventPublisher,
	)
}

func (suite *ProgressionManagementTestSuite) checkMockAssertions() {
	suite.progressionFinder.AssertExpectations(suite.T())
	suite.progressionPersistence.AssertExpectations(suite.T())
	suite.adventurerFinder.AssertExpectations(suite.T())
	suite.adventurerService.AssertExpectations(suite.T())
	suite.eventPublisher.AssertExpectations(suite.T())
}

func (suite *ProgressionManagementTestSuite) TestCreateInitialLevelShouldWork() {

	command := leveling.CreateLevel{
		NewTopExp: 20,
	}

	expectedResponse := leveling.ExpLevel{
		ID:       uuid.New(),
		Level:    1,
		StartExp: 1,
		EndExp:   20,
	}

	suite.progressionFinder.On("FindLatestLevel").Return(leveling.ExpLevel{}, nil)
	suite.progressionPersistence.On("Save", mock.AnythingOfType("leveling.ExpLevel")).Return(expectedResponse, nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(leveling.ExpLevel)
		assert.Equal(suite.T(), 1, arg.Level)
		assert.Equal(suite.T(), int64(1), arg.StartExp)
		assert.Equal(suite.T(), int64(20), arg.EndExp)
	})

	expLevel, err := suite.progressionManagement.CreateLevel(command)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedResponse, expLevel)

	suite.checkMockAssertions()
}

func (suite *ProgressionManagementTestSuite) TestAddNewLevelShouldBeAfterLatestLevel() {

	command := leveling.CreateLevel{
		NewTopExp: 40,
	}

	latestExistingLevel := leveling.ExpLevel{
		ID:       uuid.New(),
		Level:    1,
		StartExp: 1,
		EndExp:   20,
	}

	expectedResponse := leveling.ExpLevel{
		ID:       uuid.New(),
		Level:    2,
		StartExp: 21,
		EndExp:   40,
	}

	suite.progressionFinder.On("FindLatestLevel").Return(latestExistingLevel, nil)
	suite.progressionPersistence.On("Save", mock.AnythingOfType("leveling.ExpLevel")).Return(expectedResponse, nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(leveling.ExpLevel)
		assert.Equal(suite.T(), 2, arg.Level)
		assert.Equal(suite.T(), int64(21), arg.StartExp)
		assert.Equal(suite.T(), int64(40), arg.EndExp)
	})

	expLevel, err := suite.progressionManagement.CreateLevel(command)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedResponse, expLevel)

	suite.checkMockAssertions()
}

func (suite *ProgressionManagementTestSuite) TestLevelFromExperienceLargerThanToExperience() {
	command := leveling.CreateLevel{
		NewTopExp: 10,
	}

	latestExistingLevel := leveling.ExpLevel{
		ID:       uuid.New(),
		Level:    1,
		StartExp: 1,
		EndExp:   20,
	}

	suite.progressionFinder.On("FindLatestLevel").Return(latestExistingLevel, nil)

	_, err := suite.progressionManagement.CreateLevel(command)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), leveling.ErrNewLevelCantBeLowerThanLargestExistingLevel.Message, err.Error())

	suite.checkMockAssertions()
}

func (suite *ProgressionManagementTestSuite) TestAwardedExperienceLevelsCharacterUp() {
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

	suite.progressionFinder.On("FindLevelForExperience", int64(20)).Return(leveling.ExpLevel{
		ID:       uuid.UUID{},
		Level:    2,
		StartExp: 15,
		EndExp:   30,
	}, nil)

	suite.eventPublisher.On("Publish", leveling.AdventurerLevelingTopic, leveling.AdventurerLevelingEvent{
		AdventurerID: adventurerId,
		OldLevel:     1,
		NewLevel:     2,
	}).Return(leveling.AdventurerLevelingEvent{}, nil)

	suite.adventurerService.On("UpdateCharacter", adventurerId, expectedCharacter).Return(adventurers.Adventurer{}, nil)

	err := suite.progressionManagement.AwardExperience(command)

	assert.Nil(suite.T(), err)

	suite.checkMockAssertions()
}

func (suite *ProgressionManagementTestSuite) TestAwardedExperienceDoesNotLevelsCharacterUp() {
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
	suite.progressionFinder.On("FindLevelForExperience", int64(11)).Return(leveling.ExpLevel{
		ID:       uuid.UUID{},
		Level:    1,
		StartExp: 10,
		EndExp:   20,
	}, nil)

	suite.adventurerService.On("UpdateCharacter", adventurerId, expectedCharacter).Return(adventurers.Adventurer{}, nil)

	err := suite.progressionManagement.AwardExperience(command)

	assert.Nil(suite.T(), err)

	suite.checkMockAssertions()
}

func TestRunProgressionManagementSuite(t *testing.T) {
	suite.Run(t, new(ProgressionManagementTestSuite))
}
