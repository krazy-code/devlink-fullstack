package configs

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(env *Env) *gorm.DB {
	conn := NewDBConnection(DBConnectionConfig{
		Driver:          SQLDriverNamePostgres,
		DSN:             env.PostgresDSN,
		MaxIdleConns:    env.PostgresMaxIdleConns,
		MaxOpenConns:    env.PostgresMaxOpenConns,
		ConnMaxLifetime: time.Duration(env.PostgresConnMaxLifetime) * time.Second,
		ConnMaxIdleTime: time.Duration(env.PostgresConnMaxIdleTime) * time.Second,
	})

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
		// PreferSimpleProtocol: true,
	}), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}
