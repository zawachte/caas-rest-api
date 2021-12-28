package caasdb

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	dbPool *pgxpool.Pool
}

func NewClient() (*Client, error) {
	// ensure to change values as needed.
	databaseUrl := "postgres://zach:zach@localhost:5400/zach"
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	return &Client{dbPool}, nil
}
