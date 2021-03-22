package security

import (
	"github.com/orlangure/gnomock"
	mongoPreset "github.com/orlangure/gnomock/preset/mongo"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"tasquest.com/server/security"
	"tasquest.com/tests/utils"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	dbClient  *mongo.Database
	container *gnomock.Container
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	preset := mongoPreset.Preset(mongoPreset.WithUser("tasquest", "tasquest"))
	container, err := gnomock.Start(preset)

	if err != nil {
		panic(err)
	}

	suite.container = container
	suite.dbClient = utils.SetupMongo(container).Database("tasquest")
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	userRepository := security.ProvideMongoUserRepository(suite.dbClient)

	userToCreate := security.User{
		Email:     "a@a.com",
		Password:  "test",
		Active:    false,
		Providers: nil,
	}

	createdUser, err := userRepository.Save(userToCreate)

	suite.Nil(err)
	suite.NotNil(createdUser)
	suite.Equal("a@a.com", createdUser.Email)
	suite.Equal("test", createdUser.Password)
	suite.Equal(false, createdUser.Active)
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	_ = gnomock.Stop(suite.container)
}

func TestRunUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
