package gormit

import (
	fmwgorm "gorm.io/gorm"

	"tasquest.com/server/adapters/output/databases/gorm"
	"tasquest.com/server/infra/database"
)

func ITPostgresInstance() *fmwgorm.DB {
	// Fixed for now, i'll get from dynamic sources later, i'm really lazy at the moment
	config := &database.DbConfig{
		Hostname: "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
		SslMode:  "disable",
		TimeZone: "Europe/Amsterdam",
	}

	return gorm.NewGormPostgres(config)
}
