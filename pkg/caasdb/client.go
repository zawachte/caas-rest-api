package caasdb

import (
	"context"
	"strconv"
	"time"

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

func (cs *CaasServer) postCluster(cluster Cluster) error {
	var lastInsertID int

	i64, err := strconv.ParseInt(*cluster.AccountId, 10, 32)
	if err != nil {
		return err
	}

	err = cs.dbPool.QueryRow(context.Background(), "INSERT INTO clusters(account_id, kubeconfig, created_on) VALUES($1, $2, $3) returning cluster_id;",
		i64,
		cluster.Kubeconfig,
		time.Now()).Scan(&lastInsertID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CaasServer) postAccount(account Account) error {
	var lastInsertID int

	err := cs.dbPool.QueryRow(context.Background(), "INSERT INTO accounts(username, password, email, created_on) VALUES($1, $2, $3, $4) returning account_id;",
		account.Username,
		account.Password,
		account.Email,
		time.Now()).Scan(&lastInsertID)
	if err != nil {
		return err
	}

	return nil
}
