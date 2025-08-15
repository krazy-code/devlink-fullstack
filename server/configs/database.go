package configs

import (
	"database/sql"
	"time"
)

const (
	SQLDriverNamePostgres = "pgx"
	SQLDriverNameMySQL    = "mysql"
)

type DBConnectionConfig struct {
	// Required.
	Driver string
	// Required.
	DSN string
	// Default: 10
	MaxIdleConns int
	// Default: 10
	MaxOpenConns int
	// Default: 180 seconds
	ConnMaxLifetime time.Duration
	// Default: 60 seconds
	ConnMaxIdleTime time.Duration

	// Default: false
	Debug bool
}

func NewDBConnection(config DBConnectionConfig) *sql.DB {
	conn, err := sql.Open(config.Driver, config.DSN)
	if err != nil {
		panic(err)
	}

	if config.MaxIdleConns <= 0 {
		config.MaxIdleConns = 10
	}
	if config.MaxOpenConns <= 0 {
		config.MaxOpenConns = 10
	}
	if config.ConnMaxLifetime <= 0 {
		config.ConnMaxLifetime = 180 * time.Second
	}
	if config.ConnMaxIdleTime <= 0 {
		config.ConnMaxIdleTime = 60 * time.Second
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	conn.SetMaxIdleConns(config.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	conn.SetMaxOpenConns(config.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	conn.SetConnMaxLifetime(config.ConnMaxLifetime)
	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	conn.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	return conn
}
