//+build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"tasquest.com/server/adapters/output/databases/mongorepositories"
	"tasquest.com/server/api"
	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/application/gamification/tasks"
	"tasquest.com/server/application/security"
	"tasquest.com/server/commons"
	"tasquest.com/server/infra/database"
	"tasquest.com/server/infra/web"
)

/**************************
 *    Common Providers    *
 **************************/
func loggerWire() *logrus.Logger {
	panic(
		wire.Build(commons.ProvideLogger),
	)
}

/**************************
 *    Infra Providers     *
 **************************/
func databaseWire() *mongo.Database {
	panic(
		wire.Build(database.ProvideDatasource),
	)
}

func webServerWire() *gin.Engine {
	panic(
		wire.Build(web.ProvideWebServer),
	)
}

/**************************
 *    security Providers  *
 **************************/
func userFinderWire() security.UserFinder {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewUserRepository,
			wire.Bind(new(security.UserFinder), new(*mongorepositories.MongoUserRepository)),
		),
	)
}

func userPersistenceWire() security.UserPersistence {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewUserRepository,
			wire.Bind(new(security.UserPersistence), new(*mongorepositories.MongoUserRepository)),
		),
	)
}

func userServiceWire() security.UserService {
	panic(
		wire.Build(
			userPersistenceWire,
			userFinderWire,
			security.NewUserManagement,
			wire.Bind(new(security.UserService), new(*security.UserManagement)),
		),
	)
}

/*****************************
 *    Adventurers Providers  *
 *****************************/
func adventurerFinderWire() adventurers.AdventurerFinder {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewMongoAdventurerRepository,
			wire.Bind(new(adventurers.AdventurerFinder), new(*mongorepositories.MongoAdventurerRepository)),
		),
	)
}

func adventurerPersistenceWire() adventurers.AdventurerPersistence {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewMongoAdventurerRepository,
			wire.Bind(new(adventurers.AdventurerPersistence), new(*mongorepositories.MongoAdventurerRepository)),
		),
	)
}

func adventurerServiceWire() adventurers.AdventurerService {
	panic(
		wire.Build(
			userServiceWire,
			adventurerFinderWire,
			adventurerPersistenceWire,
			adventurers.NewAdventurerManager,
			wire.Bind(new(adventurers.AdventurerService), new(*adventurers.AdventurerManager)),
		),
	)
}

/*****************************
 *      Task Providers       *
 *****************************/
func taskFinderWire() tasks.TaskFinder {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.ProvideMongoTaskRepository,
			wire.Bind(new(tasks.TaskFinder), new(*mongorepositories.MongoTaskRepository)),
		),
	)
}

func taskPersistenceWire() tasks.TaskPersistence {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.ProvideMongoTaskRepository,
			wire.Bind(new(tasks.TaskPersistence), new(*mongorepositories.MongoTaskRepository)),
		),
	)
}

func taskServiceWire() tasks.TaskService {
	panic(
		wire.Build(
			taskFinderWire,
			taskPersistenceWire,
			adventurerServiceWire,
			tasks.NewTaskManager,
			wire.Bind(new(tasks.TaskService), new(*tasks.TaskManager)),
		),
	)
}

/**************************
 *    Api Providers  *
 **************************/
func authApiWireBuilder() *api.AuthAPI {
	panic(
		wire.Build(
			web.ProvideWebServer,
			commons.ProvideLogger,
			commons.ProvideErrorHandler,
			userServiceWire,
			api.ProvideAuthAPI,
		),
	)
}

// ########################## Initializer #########################

func Bootstrap() {
	// Common
	loggerWire()
	// Infra
	databaseWire()
	webServerWire()
	// security
	userPersistenceWire()
	userFinderWire()
	userServiceWire()
	// Adventurer
	adventurerFinderWire()
	adventurerPersistenceWire()
	adventurerServiceWire()
	// Tasks
	taskFinderWire()
	taskPersistenceWire()
	taskServiceWire()
	// Apis
	authApiWireBuilder()
}
