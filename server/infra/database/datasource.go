package database

import (
	"context"
	"github.com/kkyr/fig"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"sync"
	"tasquest.com/server/common"
	"time"
)

type DBConfig struct {
	Host     string `fig:"host" default:"mongodb://localhost:27017"`
	Timeout  int    `fig:"timeout" default:"5"`
	User     string `fig:"user" default:"tasquest"`
	Password string `fig:"password" default:"tasquest"`
}

var once sync.Once
var instance *mongo.Client

func ProvideDatasource() *mongo.Client {
	once.Do(func() {
		instance = doConnect(buildOptions())
	})
	return instance
}

func doConnect(options *options.ClientOptions) *mongo.Client {
	log.Info("Connecting to the Database at " + strings.Join(options.Hosts, ","))

	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		common.IrrecoverableFailure("Failed to connect to the database!", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		common.IrrecoverableFailure("Unreachable database!", err)
	}

	log.Info("Database connected at " + strings.Join(options.Hosts, ","))

	return client
}

func buildOptions() *options.ClientOptions {
	config := readConfig()

	credentials := options.Credential{
		Username: config.User,
		Password: config.Password,
	}

	return options.Client().ApplyURI(config.Host).SetAppName("TasQuest").SetAuth(credentials).SetConnectTimeout(time.Duration(config.Timeout) * time.Second)
}

func readConfig() DBConfig {
	var config DBConfig

	log.Info("Loading Database configurations")

	err := fig.Load(
		&config,
		fig.File("dbconfig.yml"),
		fig.Dirs("./config"),
	)

	if err != nil {
		common.IrrecoverableFailure("Failed to fetch the database config file", err)
	}

	return config
}
