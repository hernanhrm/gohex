package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func LoadDatabaseConnection(config LocalConfig) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		config.Database.Driver,
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	))
	if err != nil {
		return nil, fmt.Errorf("Unable to create connection pool: %v\n", err)
	}
	defer dbPool.Close()

	if err := dbPool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("Unable to ping database: %v\n", err)
	}

	return dbPool, nil
}
