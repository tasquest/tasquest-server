// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"tasquest.com/server/adapters/input/rest"
	"tasquest.com/server/adapters/output/databases/mongorepositories"
	message_bus "tasquest.com/server/adapters/output/eventbus/message-bus"
	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/application/gamification/leveling"
	"tasquest.com/server/application/gamification/tasks"
	"tasquest.com/server/application/security"
	"tasquest.com/server/commons"
	"tasquest.com/server/infra/events"
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
		wire.Build(mongorepositories.ProvideDatasource),
	)
}

func webServerWire() *gin.Engine {
	panic(
		wire.Build(web.ProvideWebServer),
	)
}

func eventPublisherWire() events.Publisher {
	panic(
		wire.Build(
			message_bus.NewMessenger,
			wire.Bind(new(events.Publisher), new(*message_bus.Messenger)),
		),
	)
}

func eventSubscriberWire() events.Subscriber {
	panic(
		wire.Build(
			message_bus.NewMessenger,
			wire.Bind(new(events.Subscriber), new(*message_bus.Messenger)),
		),
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
			wire.Bind(new(security.UserFinder), new(*mongorepositories.UserRepository)),
		),
	)
}

func userPersistenceWire() security.UserPersistence {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewUserRepository,
			wire.Bind(new(security.UserPersistence), new(*mongorepositories.UserRepository)),
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
			mongorepositories.NewAdventurerRepository,
			wire.Bind(new(adventurers.AdventurerFinder), new(*mongorepositories.AdventurerRepository)),
		),
	)
}

func adventurerPersistenceWire() adventurers.AdventurerPersistence {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewAdventurerRepository,
			wire.Bind(new(adventurers.AdventurerPersistence), new(*mongorepositories.AdventurerRepository)),
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
			mongorepositories.NewTaskRepository,
			wire.Bind(new(tasks.TaskFinder), new(*mongorepositories.TaskRepository)),
		),
	)
}

func taskPersistenceWire() tasks.TaskPersistence {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewTaskRepository,
			wire.Bind(new(tasks.TaskPersistence), new(*mongorepositories.TaskRepository)),
		),
	)
}

func taskServiceWire() tasks.TaskService {
	panic(
		wire.Build(
			taskFinderWire,
			taskPersistenceWire,
			adventurerFinderWire,
			tasks.NewTaskManager,
			wire.Bind(new(tasks.TaskService), new(*tasks.TaskManager)),
		),
	)
}

/*****************************
 *   Progression Providers   *
 *****************************/

func progressionFinderWire() leveling.ProgressionFinder {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewProgressionRepository,
			wire.Bind(new(leveling.ProgressionFinder), new(*mongorepositories.ProgressionRepository)),
		),
	)
}

func progressionPersistenceWire() leveling.ProgressionPersistence {
	panic(
		wire.Build(
			databaseWire,
			mongorepositories.NewProgressionRepository,
			wire.Bind(new(leveling.ProgressionPersistence), new(*mongorepositories.ProgressionRepository)),
		),
	)
}

func progressionServiceWire() leveling.ProgressionService {
	panic(
		wire.Build(
			progressionFinderWire,
			progressionPersistenceWire,
			adventurerServiceWire,
			adventurerFinderWire,
			eventPublisherWire,
			leveling.NewProgressionManagement,
			wire.Bind(new(leveling.ProgressionService), new(*leveling.ProgressionManagement)),
		),
	)
}

/**************************
 *    Api Providers  *
 **************************/
func authApiWireBuilder() *rest.AuthAPI {
	panic(
		wire.Build(
			web.ProvideWebServer,
			commons.ProvideLogger,
			commons.ProvideErrorHandler,
			userServiceWire,
			rest.NewAuthApi,
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
	eventPublisherWire()
	eventSubscriberWire()
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
	// Leveling
	progressionFinderWire()
	progressionPersistenceWire()
	progressionServiceWire()
	// Apis
	authApiWireBuilder()
}
