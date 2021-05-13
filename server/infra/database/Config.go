package database

import (
	"github.com/kkyr/fig"
	log "github.com/sirupsen/logrus"

	"tasquest.com/server/commons"
)

type DbConfig struct {
	Hostname string `fig:"hostname" default:"localhost"`
	Port     int    `fig:"port" default:"9920"`
	User     string `fig:"user" default:"tasquest"`
	Password string `fig:"password" default:"tasquest"`
	Dbname   string `fig:"dbname" default:"tasquest"`
	SslMode  string `fig:"sslMode" default:"disable"`
	TimeZone string `fig:"timezone"`
}

func ReadFromDisk() *DbConfig {
	var config DbConfig

	log.Info("Loading Database configurations")

	err := fig.Load(
		&config,
		fig.File("dbconfig.yml"),
		fig.Dirs("./config"),
	)

	if err != nil {
		commons.IrrecoverableFailure("Failed to fetch the database config file", err)
	}

	return &config
}
