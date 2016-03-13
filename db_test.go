package kkreport

import (
	"testing"

	"github.com/jackc/pgx"
	"github.com/satori/go.uuid"
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

func testDBMethods(t *testing.T) {
	testInvalidInsert(t)

	testInsert(t)
	testGetAndHandle(t)

	truncate(t)
}

func testInvalidInsert(t *testing.T) {
	var one Report
	one.ID = "abc"
	if err := insertReport(&one); err == nil {
		t.Error("Should have error with invalid ID.")
	}
}

func testInsert(t *testing.T) {
	var one Report
	one.ID = uuid.NewV1().String()
	one.UserID = 3
	one.Reason = 0
	one.Value = "abc"

	if err := insertReport(&one); err != nil {
		t.Fatal(err)
	}

	one.ID = uuid.NewV1().String()
	if err := insertReport(&one); err != nil {
		t.Fatal(err)
	}

	one.ID = uuid.NewV1().String()
	if err := insertReport(&one); err != nil {
		t.Fatal(err)
	}
}

func testGetAndHandle(t *testing.T) {
	reports, err := getAll(0)
	if err != nil {
		t.Error(err)
	} else if len(reports) != 3 {
		t.Error("Result is wrong.")
	}

	// mark one as handled
	if err := handleOne(reports[0].ID); err != nil {
		t.Error(err)
	}

	// get unhandled reports
	if unhandledOnes, err := getUnHandled(0); err != nil {
		t.Error(err)
	} else if len(unhandledOnes) != 2 {
		t.Error("Result is wrong.")
	}

	// get handled reports
	if handledOnes, err := getHandled(0); err != nil {
		t.Error(err)
	} else if len(handledOnes) != 1 {
		t.Error("Result is wrong.")
	}

	if err := deleteOne(reports[0].ID); err != nil {
		t.Error(err)
	}

	if all, err := getAll(0); err != nil {
		t.Error(err)
	} else if len(all) != 2 {
		t.Error("Result is wrong.")
	}
}

func truncate(t *testing.T) {
	if _, err := dbPool.Exec("TRUNCATE report"); err != nil {
		t.Error(err)
	}
}
