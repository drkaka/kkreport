package kkreport

import "github.com/jackc/pgx"

var dbPool *pgx.ConnPool

// prepareDB to prepare the database.
func prepareDB() error {
	s := `CREATE TABLE IF NOT EXISTS report (
	id uuid primary key,
	userid integer,
    handle boolean DEFAULT false,
    reason smallint,
    value text);`

	_, err := dbPool.Exec(s)
	return err
}
