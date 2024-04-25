package postgres

import (
	"fmt"
	"github.com/zer0day88/brick-test/pkg/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {

	var (
		host            = config.Key.Database.Postgres.Host
		user            = config.Key.Database.Postgres.User
		password        = config.Key.Database.Postgres.Password
		port            = config.Key.Database.Postgres.Port
		dbname          = config.Key.Database.Postgres.DbName
		connMaxLifetime = config.Key.Database.Postgres.ConnectionMaxLifetime
		connMaxOpen     = config.Key.Database.Postgres.ConnectionMaxOpen
		connMaxIdle     = config.Key.Database.Postgres.ConnectionMaxIdle
		connMaxIdleTime = config.Key.Database.Postgres.ConnectionMaxIdleTime
		sslMode         = config.Key.Database.Postgres.SSLMode
	)

	dsn := fmt.Sprintf("host=%s user='%s' password='%s' dbname=%s port=%d sslmode=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(connMaxIdleTime) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	sqlDB.SetMaxOpenConns(connMaxOpen)
	sqlDB.SetMaxIdleConns(connMaxIdle)

	return db, nil
}
