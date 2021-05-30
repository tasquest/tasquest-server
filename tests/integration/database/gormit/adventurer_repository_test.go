package gormit

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	fmwgorm "gorm.io/gorm"

	"tasquest.com/server/adapters/output/databases/gorm"
	"tasquest.com/tests/fixtures"
)

type AdventurerRepositoryIntegrationTestSuite struct {
	suite.Suite
	// Dependencies
	db *fmwgorm.DB
	// InjectTo
	adventurerRepository *gorm.AdventurersGormRepository
}

func (suite *AdventurerRepositoryIntegrationTestSuite) SetupTest() {
	gorm.IsProgressionGormRepositoryInstanced = sync.Once{}
	suite.db = ITPostgresInstance().Begin()
	suite.adventurerRepository = gorm.NewAdventurersGormRepository(suite.db)
}

func (suite *AdventurerRepositoryIntegrationTestSuite) TestSaveUpdateFetch() {
	adventurer := fixtures.NewAdventurer()

	_, err := suite.adventurerRepository.Save(adventurer)
	assert.Nil(suite.T(), err)

	savedAdventurer, err := suite.adventurerRepository.FindByID(adventurer.ID)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), adventurer.Name, savedAdventurer.Name)
	assert.Equal(suite.T(), adventurer.Surname, savedAdventurer.Surname)
	assert.Equal(suite.T(), adventurer.ID, savedAdventurer.Character.AdventurerID)
	assert.Equal(suite.T(), adventurer.Character.Level, savedAdventurer.Character.Level)
	assert.Equal(suite.T(), adventurer.Character.Title, savedAdventurer.Character.Title)
	assert.Equal(suite.T(), adventurer.Character.Experience, savedAdventurer.Character.Experience)

	savedAdventurer.Name = "Jack"

	updatedAdventurer, err := suite.adventurerRepository.Update(savedAdventurer)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Jack", updatedAdventurer.Name)

	suite.db.Rollback()
}

func TestRunAdventurerRepositoryIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(AdventurerRepositoryIntegrationTestSuite))
}
