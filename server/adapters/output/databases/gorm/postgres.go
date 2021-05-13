package gorm

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"tasquest.com/server/commons"
	"tasquest.com/server/infra/database"
)

var IsGormPostgresInstanced sync.Once
var gormPostgresInstance *gorm.DB

func NewGormPostgres(config *database.DbConfig) *gorm.DB {
	IsGormPostgresInstanced.Do(func() {
		db, err := gorm.Open(postgres.New(
			postgres.Config{DSN: parseDsn(*config)},
		))

		if err != nil {
			commons.IrrecoverableFailure("Failed to connect to the database", err)
		}

		gormPostgresInstance = db
	})

	return gormPostgresInstance
}

func parseDsn(config database.DbConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Hostname,
		config.User,
		config.Password,
		config.Dbname,
		config.Port,
		config.SslMode,
		config.TimeZone,
	)
}
