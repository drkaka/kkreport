package kkreport

import "github.com/jackc/pgx"

const (
	insert    = "INSERT INTO report(id,userid,at,reason,value) VALUES($1,$2,$3,$4,$5)"
	all       = "SELECT id,userid,at,handle,reason,value FROM report WHERE at>=$1"
	handled   = "SELECT id,userid,at,handle,reason,value FROM report WHERE at>=$1 AND handle=true"
	unhandled = "SELECT id,userid,at,handle,reason,value FROM report WHERE at>=$1 AND handle=false"
)

var dbPool *pgx.ConnPool

// prepareDB to prepare the database.
func prepareDB() error {
	s := `CREATE TABLE IF NOT EXISTS report (
	id uuid primary key,
	userid integer,
    at integer,
    handle boolean DEFAULT false,
    reason smallint,
    value text);`

	_, err := dbPool.Exec(s)
	return err
}

// insertReport to insert a report to database.
func insertReport(repo *Report) error {
	_, err := dbPool.Exec(insert, repo.ID, repo.UserID, repo.At, repo.Reason, repo.Value)
	return err
}

// handleOne to mark a report as handled
func handleOne(id string) error {
	_, err := dbPool.Exec("UPDATE report SET handle=true WHERE id=$1", id)
	return err
}

// deleteOne to delete one report
func deleteOne(id string) error {
	_, err := dbPool.Exec("DELETE FROM report WHERE id=$1", id)
	return err
}

// getAll to get all the reports.
// utime the unixtime, the repots will be got after that time.
func getAll(utime int32) ([]Report, error) {
	rows, _ := dbPool.Query(all, utime)

	var result []Report
	for rows.Next() {
		var one Report
		err := rows.Scan(&(one.ID), &(one.UserID), &(one.At), &(one.Handle), &(one.Reason), &(one.Value))
		if err != nil {
			return result, err
		}
		result = append(result, one)
	}

	return result, rows.Err()
}

// getUnHandled to get unhandled the reports.
// utime the unixtime, the repots will be got after that time.
func getUnHandled(utime int32) ([]Report, error) {
	rows, _ := dbPool.Query(unhandled, utime)

	var result []Report
	for rows.Next() {
		var one Report
		err := rows.Scan(&(one.ID), &(one.UserID), &(one.At), &(one.Handle), &(one.Reason), &(one.Value))
		if err != nil {
			return result, err
		}
		result = append(result, one)
	}

	return result, rows.Err()
}

// getHandled to get handled the reports.
// utime the unixtime, the repots will be got after that time.
func getHandled(utime int32) ([]Report, error) {
	rows, _ := dbPool.Query(handled, utime)

	var result []Report
	for rows.Next() {
		var one Report
		err := rows.Scan(&(one.ID), &(one.UserID), &(one.At), &(one.Handle), &(one.Reason), &(one.Value))
		if err != nil {
			return result, err
		}
		result = append(result, one)
	}

	return result, rows.Err()
}
