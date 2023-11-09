package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(databaseUrl string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
