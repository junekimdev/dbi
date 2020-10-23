package dbi

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

// Connect to DB and load db-pool reference in the package
func Connect(connString string) error {
	// Connect to Postgres
	pgpool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return err
	}
	db = pgpool
	return nil
}

// Query is the wrapper function of 'pgxpool.Query'
//
// This starts a goroutine to run query and return the result (pgx.Rows) through the go-channel
func Query(sql string, args ...interface{}) pgx.Rows {
	resChan := make(chan pgx.Rows)

	go func() {
		rows, err := db.Query(context.Background(), sql, args...)
		if err != nil {
			log.Fatalf("Unable to query the database: %v", err)
			resChan <- nil
		}
		resChan <- rows
	}()
	return <-resChan
}

// Scan is the wrapper function of 'pgx.Rows.Scan' that scan the result of query
//
// You Need to pass "scan function" that has "Scan" method of pgx.Rows
func Scan(rows pgx.Rows, scanfunc func()) error {
	defer rows.Close()

	for rows.Next() {
		// Pass all pointers of dest variable
		scanfunc()
	}

	if rows.Err() != nil {
		log.Fatalf("Failed to scan the query: %v\n", rows.Err())
		return rows.Err()
	}

	return nil
}
