// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"

	fmw_gorm "gorm.io/gorm"

	"tasquest.com/server/adapters/input/rest"
	"tasquest.com/server/adapters/output/databases/gorm"
	message_bus "tasquest.com/server/adapters/output/eventbus/message-bus"
	"tasquest.com/server/application"
	"tasquest.com/server/application/gamification/adventurers"
	"tasquest.com/server/application/gamification/leveling"
	"tasquest.com/server/application/gamification/tasks"
	"tasquest.com/server/application/security"
	"tasquest.com/server/commons"
	"tasquest.com/server/infra/database"
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
func databaseWire() *fmw_gorm.DB {
	panic(
		wire.Build(
			database.ReadFromDisk,
			gorm.NewGormPostgres,
		),
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

func subscriberManagerWire() *application.SubscriptionManagement {
	panic(
		wire.Build(
			levelingSubscriberWire,
			application.NewSubscriptionManagement,
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
			gorm.NewUserGormRepository,
			wire.Bind(new(security.UserFinder), new(*gorm.UsersGormRepository)),
		),
	)
}

func userPersistenceWire() security.UserPersistence {
	panic(
		wire.Build(
			databaseWire,
			gorm.NewUserGormRepository,
			wire.Bind(new(security.UserPersistence), new(*gorm.UsersGormRepository)),
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
			gorm.NewAdventurersGormRepository,
			wire.Bind(new(adventurers.AdventurerFinder), new(*gorm.AdventurersGormRepository)),
		),
	)
}

func adventurerPersistenceWire() adventurers.AdventurerPersistence {
	panic(
		wire.Build(
			databaseWire,
			gorm.NewAdventurersGormRepository,
			wire.Bind(new(adventurers.AdventurerPersistence), new(*gorm.AdventurersGormRepository)),
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
			gorm.NewTasksGormRepository,
			wire.Bind(new(tasks.TaskFinder), new(*gorm.TasksGormRepository)),
		),
	)
}

func taskPersistenceWire() tasks.TaskPersistence {
	panic(
		wire.Build(
			databaseWire,
			gorm.NewTasksGormRepository,
			wire.Bind(new(tasks.TaskPersistence), new(*gorm.TasksGormRepository)),
		),
	)
}

func taskServiceWire() tasks.TaskService {
	panic(
		wire.Build(
			taskFinderWire,
			taskPersistenceWire,
			adventurerFinderWire,
			eventPublisherWire,
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
			gorm.NewProgressionGormRepository,
			wire.Bind(new(leveling.ProgressionFinder), new(*gorm.ProgressionGormRepository)),
		),
	)
}

func progressionPersistenceWire() leveling.ProgressionPersistence {
	panic(
		wire.Build(
			databaseWire,
			gorm.NewProgressionGormRepository,
			wire.Bind(new(leveling.ProgressionPersistence), new(*gorm.ProgressionGormRepository)),
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

func levelingSubscriberWire() *leveling.Subscribers {
	panic(
		wire.Build(
			eventSubscriberWire,
			progressionServiceWire,
			leveling.NewLevelingSubscribers,
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
	subscriberManagerWire()
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
	levelingSubscriberWire()
	// Apis
	authApiWireBuilder()
}
