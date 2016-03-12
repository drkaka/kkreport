package kkreport

import (
	"testing"

	"github.com/jackc/pgx"
)

func testTableGeneration(t *testing.T) {
	var dbname pgx.NullString
	if err := dbPool.QueryRow("SELECT 'public.report'::regclass;").Scan(&dbname); err != nil {
		t.Fatal(err)
	}

	if dbname.String != "report" {
		t.Fatal("dbname is not correct.")
	}
}
