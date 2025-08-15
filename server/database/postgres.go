package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krazy-code/devlink/utils"
)

func PostgreSQLConnection() (*pgxpool.Pool, error) {
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	postgresConnURL, err := utils.ConnectionURLBuilder(os.Getenv("DATABASE_NAME"))
	if err != nil {
		return nil, err
	}

	config, err := pgxpool.ParseConfig(postgresConnURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing database URL: %w", err)
	}
	config.MaxConns = int32(maxConn)
	config.MinConns = int32(maxIdleConn)
	config.MaxConnLifetime = time.Duration(maxLifetimeConn) * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %w", err)
	}
	return pool, nil
}
