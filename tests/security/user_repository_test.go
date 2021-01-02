package security

import (
	"context"
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
	dbClient  *mongo.Client
	container *gnomock.Container
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	preset := mongoPreset.Preset(mongoPreset.WithUser("tasquest", "tasquest"))
	container, err := gnomock.Start(preset)

	if err != nil {
		panic(err)
	}

	suite.container = container
	suite.dbClient = utils.SetupMongo(container)
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	userRepository := security.MongoUserRepository{
		Collection: suite.dbClient.Database("tasquest").Collection("users"),
	}

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
	suite.dbClient.Disconnect(context.TODO())
}

func TestRunUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
