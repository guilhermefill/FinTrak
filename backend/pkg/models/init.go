package models

import (
	"context"
	"log"

	"github.com/guilhermefill/FinTrak/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbHost   = utils.GetEnvVars("PGHOST")
	dbPort   = utils.GetEnvVars("PGPORT")
	dbName   = utils.GetEnvVars("PGDATABASE")
	dbUser   = utils.GetEnvVars("PGUSER")
	dbPass   = utils.GetEnvVars("PGPASS")
	DbClient *pgxpool.Pool
)

func NewPool(c context.Context) (*pgxpool.Pool, error) {
	connStr := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName
	pool, err := pgxpool.New(c, connStr)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func Init(c context.Context) error {
	var err error
	DbClient, err = NewPool(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established with: " + dbName)
	return nil
}
