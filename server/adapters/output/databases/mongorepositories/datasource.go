package mongorepositories

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/kkyr/fig"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"tasquest.com/server/commons"
)

type DBConfig struct {
	Host     string `fig:"host" default:"mongodb://localhost:27017"`
	Timeout  int    `fig:"timeout" default:"5"`
	User     string `fig:"user" default:"tasquest"`
	Password string `fig:"password" default:"tasquest"`
}

var once sync.Once
var instance *mongo.Database

func ProvideDatasource() *mongo.Database {
	once.Do(func() {
		instance = doConnect(buildOptions()).Database("tasquest")
	})
	return instance
}

func doConnect(options *options.ClientOptions) *mongo.Client {
	log.Info("Connecting to the Database at " + strings.Join(options.Hosts, ","))

	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		commons.IrrecoverableFailure("Failed to connect to the database!", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		commons.IrrecoverableFailure("Unreachable database!", err)
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

	clientOptions := options.Client()
	clientOptions = clientOptions.ApplyURI(config.Host)
	clientOptions = clientOptions.SetAppName("TasQuest")
	clientOptions = clientOptions.SetAuth(credentials)
	clientOptions = clientOptions.SetConnectTimeout(time.Duration(config.Timeout) * time.Second)
	return clientOptions
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
		commons.IrrecoverableFailure("Failed to fetch the database config file", err)
	}

	return config
}
