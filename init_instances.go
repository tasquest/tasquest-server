//+build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"tasquest.com/server/api"
	"tasquest.com/server/common/errorhandler"
	"tasquest.com/server/common/logs"
	"tasquest.com/server/domain/profiles"
	"tasquest.com/server/infra/database"
	"tasquest.com/server/infra/web"
	"tasquest.com/server/security"
)

/**************************
 *    Common Providers    *
 **************************/
func loggerWireBuilder() *logrus.Logger {
	wire.Build(logs.ProvideLogger)
	return &logrus.Logger{}
}

/**************************
 *    Infra Providers     *
 **************************/
func databaseWireBuilder() *mongo.Client {
	wire.Build(database.ProvideDatasource)
	return &mongo.Client{}
}

func webServerWireBuilder() *gin.Engine {
	wire.Build(web.ProvideWebServer)
	return &gin.Engine{}
}

/**************************
 *    security Providers  *
 **************************/
func userRepositoryWireBuilder() security.UserRepository {
	panic(
		wire.Build(
			security.ProvideMongoUserRepository,
			database.ProvideDatasource,
			wire.Bind(new(security.UserRepository), new(*security.MongoUserRepository)),
		),
	)
}

func userManagementWireBuilder() security.UserManagement {
	panic(
		wire.Build(
			security.ProvideDefaultUserManagement,
			security.ProvideMongoUserRepository,
			database.ProvideDatasource,
			wire.Bind(new(security.UserRepository), new(*security.MongoUserRepository)),
			wire.Bind(new(security.UserManagement), new(*security.DefaultUserManagement)),
		),
	)
}

/**************************
 *    Profiles Providers  *
 **************************/
func profileRepositoryWireBuilder() profiles.ProfileRepository {
	panic(
		wire.Build(
			profiles.ProvideMongoProfileRepository,
			database.ProvideDatasource,
			wire.Bind(new(profiles.ProfileRepository), new(*profiles.MongoProfileRepository)),
		),
	)
}

func profileManagerWireBuilder() profiles.ProfileManager {
	panic(
		wire.Build(
			profiles.ProvideDefaultProfileManager,
			security.ProvideDefaultUserManagement,
			profiles.ProvideMongoProfileRepository,
			security.ProvideMongoUserRepository,
			database.ProvideDatasource,
			wire.Bind(new(profiles.ProfileManager), new(*profiles.DefaultProfileManager)),
			wire.Bind(new(security.UserManagement), new(*security.DefaultUserManagement)),
			wire.Bind(new(security.UserRepository), new(*security.MongoUserRepository)),
			wire.Bind(new(profiles.ProfileRepository), new(*profiles.MongoProfileRepository)),
		),
	)
}

/**************************
 *    Api Providers  *
 **************************/
func authApiWireBuilder() *api.AuthAPI {
	wire.Build(
		api.ProvideAuthAPI,
		web.ProvideWebServer,
		database.ProvideDatasource,
		errorhandler.ProvideErrorHandler,
		logs.ProvideLogger,
		security.ProvideDefaultUserManagement,
		security.ProvideMongoUserRepository,
		wire.Bind(new(security.UserManagement), new(*security.DefaultUserManagement)),
		wire.Bind(new(security.UserRepository), new(*security.MongoUserRepository)),
	)
	return &api.AuthAPI{}
}

// ########################## Initializer #########################

func Bootstrap() {
	// Common
	loggerWireBuilder()
	// Infra
	databaseWireBuilder()
	webServerWireBuilder()
	// security
	userRepositoryWireBuilder()
	userManagementWireBuilder()
	// Profile
	profileRepositoryWireBuilder()
	profileManagerWireBuilder()
	// APIs
	authApiWireBuilder()
}
