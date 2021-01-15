//+build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"tasquest.com/server"
	"tasquest.com/server/api"
	"tasquest.com/server/gamification/adventurers"
	"tasquest.com/server/gamification/tasks"
	"tasquest.com/server/infra/database"
	"tasquest.com/server/infra/web"
	"tasquest.com/server/security"
)

/**************************
 *    Common Providers    *
 **************************/
func loggerWireBuilder() *logrus.Logger {
	wire.Build(server.ProvideLogger)
	return &logrus.Logger{}
}

/**************************
 *    Infra Providers     *
 **************************/
func databaseWireBuilder() *mongo.Database {
	wire.Build(database.ProvideDatasource)
	return &mongo.Database{}
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

/*****************************
 *    Adventurers Providers  *
 *****************************/
func adventurerRepositoryWireBuilder() adventurers.AdventurerRepository {
	panic(
		wire.Build(
			adventurers.ProvideMongoAdventurerRepository,
			database.ProvideDatasource,
			wire.Bind(new(adventurers.AdventurerRepository), new(*adventurers.MongoAdventurerRepository)),
		),
	)
}

func adventurerManagerWireBuilder() adventurers.AdventurerManager {
	panic(
		wire.Build(
			adventurers.ProvideDefaultAdventurerManager,
			security.ProvideDefaultUserManagement,
			adventurers.ProvideMongoAdventurerRepository,
			security.ProvideMongoUserRepository,
			database.ProvideDatasource,
			wire.Bind(new(adventurers.AdventurerManager), new(*adventurers.DefaultAdventurerManager)),
			wire.Bind(new(security.UserManagement), new(*security.DefaultUserManagement)),
			wire.Bind(new(security.UserRepository), new(*security.MongoUserRepository)),
			wire.Bind(new(adventurers.AdventurerRepository), new(*adventurers.MongoAdventurerRepository)),
		),
	)
}

/*****************************
 *      Task Providers       *
 *****************************/
func taskRepositoryWireBuilder() tasks.TaskRepository {
	panic(
		wire.Build(
			tasks.ProvideMongoTaskRepository,
			database.ProvideDatasource,
			wire.Bind(new(tasks.TaskRepository), new(*tasks.MongoTaskRepository)),
		),
	)
}

func taskManagerWireBuilder() tasks.TaskManager {
	panic(
		wire.Build(
			tasks.ProvideMongoTaskRepository,
			tasks.ProvideDefaultTaskManager,
			adventurers.ProvideDefaultAdventurerManager,
			adventurers.ProvideMongoAdventurerRepository,
			security.ProvideDefaultUserManagement,
			security.ProvideMongoUserRepository,
			database.ProvideDatasource,
			wire.Bind(new(tasks.TaskRepository), new(*tasks.MongoTaskRepository)),
			wire.Bind(new(tasks.TaskManager), new(*tasks.TaskManagerImpl)),
			wire.Bind(new(adventurers.AdventurerManager), new(*adventurers.DefaultAdventurerManager)),
			wire.Bind(new(security.UserManagement), new(*security.DefaultUserManagement)),
			wire.Bind(new(security.UserRepository), new(*security.MongoUserRepository)),
			wire.Bind(new(adventurers.AdventurerRepository), new(*adventurers.MongoAdventurerRepository)),
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
		server.ProvideErrorHandler,
		server.ProvideLogger,
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
	// Adventurer
	adventurerRepositoryWireBuilder()
	adventurerManagerWireBuilder()
	// Tasks
	taskRepositoryWireBuilder()
	taskManagerWireBuilder()
	// APIs
	authApiWireBuilder()
}
