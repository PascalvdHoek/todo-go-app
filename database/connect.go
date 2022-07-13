package database

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func ConnectBunDb() *bun.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "docker"
		dbname   = "world"
	)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	// Open a PostgreSQL database.
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	// Create a Bun db on top of it.
	localDb := bun.NewDB(pgdb, pgdialect.New())
	// Print all queries to stdout.
	localDb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	err := localDb.Ping()

	if err != nil {
		panic(err)
		return nil
	}
	fmt.Println("Successfully connected to db")
	return localDb
}
