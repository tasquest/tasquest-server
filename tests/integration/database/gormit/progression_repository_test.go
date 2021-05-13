package gormit

import (
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	fmwgorm "gorm.io/gorm"

	"tasquest.com/server/adapters/output/databases/gorm"
	"tasquest.com/server/application/gamification/leveling"
)

type ProgressionRepositoryIntegrationTestSuite struct {
	suite.Suite
	// Dependencies
	db *fmwgorm.DB
	// InjectTo
	progressionRepository *gorm.ProgressionGormRepository
}

func (suite *ProgressionRepositoryIntegrationTestSuite) SetupTest() {
	gorm.IsProgressionGormRepositoryInstanced = sync.Once{}
	suite.db = ITPostgresInstance().Begin()
	suite.progressionRepository = gorm.NewProgressionGormRepository(suite.db)
}

func (suite *ProgressionRepositoryIntegrationTestSuite) TestSaveFetchAndDelete() {
	level := leveling.ExpLevel{ID: uuid.New(), Level: 1, StartExp: 1, EndExp: 20}

	_, err := suite.progressionRepository.Save(level)
	assert.Nil(suite.T(), err)

	savedLevel, err := suite.progressionRepository.FindByID(level.ID)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), level.Level, savedLevel.Level)
	assert.Equal(suite.T(), level.EndExp, savedLevel.EndExp)
	assert.Equal(suite.T(), level.StartExp, savedLevel.StartExp)

	_, err = suite.progressionRepository.Delete(savedLevel)

	assert.Nil(suite.T(), err)

	noLevel, err := suite.progressionRepository.FindByID(savedLevel.ID)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "record not found", err.Error())
	assert.Equal(suite.T(), leveling.ExpLevel{}, noLevel)

	suite.db.Rollback()
}

func (suite *ProgressionRepositoryIntegrationTestSuite) TestFindHighestLevel() {
	level := leveling.ExpLevel{ID: uuid.New(), Level: 1, StartExp: 1, EndExp: 20}
	level2 := leveling.ExpLevel{ID: uuid.New(), Level: 2, StartExp: 21, EndExp: 31}

	_, _ = suite.progressionRepository.Save(level)
	_, _ = suite.progressionRepository.Save(level2)

	highestLevel, err := suite.progressionRepository.FindLatestLevel()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), level2.Level, highestLevel.Level)
	assert.Equal(suite.T(), level2.EndExp, highestLevel.EndExp)
	assert.Equal(suite.T(), level2.StartExp, highestLevel.StartExp)

	suite.db.Rollback()
}

func (suite *ProgressionRepositoryIntegrationTestSuite) TestLevelForExperience() {
	level := leveling.ExpLevel{ID: uuid.New(), Level: 1, StartExp: 1, EndExp: 20}
	level2 := leveling.ExpLevel{ID: uuid.New(), Level: 2, StartExp: 21, EndExp: 30}
	level3 := leveling.ExpLevel{ID: uuid.New(), Level: 2, StartExp: 31, EndExp: 40}

	_, _ = suite.progressionRepository.Save(level)
	_, _ = suite.progressionRepository.Save(level2)
	_, _ = suite.progressionRepository.Save(level3)

	currentLevel, err := suite.progressionRepository.FindLevelForExperience(25)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), level2.Level, currentLevel.Level)
	assert.Equal(suite.T(), level2.EndExp, currentLevel.EndExp)
	assert.Equal(suite.T(), level2.StartExp, currentLevel.StartExp)

	suite.db.Rollback()
}

func TestRunProgressionRepositoryIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ProgressionRepositoryIntegrationTestSuite))
}
