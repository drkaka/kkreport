package kkreport

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx"
)

func TestMain(t *testing.T) {
	DBName := os.Getenv("dbname")
	DBHost := os.Getenv("dbhost")
	DBUser := os.Getenv("dbuser")
	DBPassword := os.Getenv("dbpassword")

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     DBHost,
			User:     DBUser,
			Password: DBPassword,
			Database: DBName,
			Dial:     (&net.Dialer{KeepAlive: 5 * time.Minute, Timeout: 5 * time.Second}).Dial,
		},
		MaxConnections: 10,
	}

	var err error
	var pool *pgx.ConnPool
	if pool, err = pgx.NewConnPool(connPoolConfig); err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	if err = Use(pool); err != nil {
		t.Fatal(err)
	}
	testTableGeneration(t)

	testDBMethods(t)
	testPubInsertReports(t)
	testPubGetAndHandle(t)

	if _, err = dbPool.Exec("DROP TABLE report"); err != nil {
		t.Error(err)
	}
}

func testPubInsertReports(t *testing.T) {
	userid := int32(3)
	reason := int16(1)
	value := "abc"

	if err := InsertReport(userid, reason, value); err != nil {
		t.Fatal(err)
	}

	if err := InsertReport(userid, reason, value); err != nil {
		t.Fatal(err)
	}

	if err := InsertReport(userid, reason, value); err != nil {
		t.Fatal(err)
	}
}

func testPubGetAndHandle(t *testing.T) {
	reports, err := GetAllReports(0)
	if err != nil {
		t.Error(err)
	} else if len(reports) != 3 {
		t.Error("Result is wrong.")
	}

	// mark one as handled
	if err := HandleReport(reports[0].ID); err != nil {
		t.Error(err)
	}

	// get unhandled reports
	if unhandledOnes, err := GetUnhandledReports(0); err != nil {
		t.Error(err)
	} else if len(unhandledOnes) != 2 {
		t.Error("Result is wrong.")
	}

	// get handled reports
	if handledOnes, err := GetHandledReports(0); err != nil {
		t.Error(err)
	} else if len(handledOnes) != 1 {
		t.Error("Result is wrong.")
	}

	if err := DeleteReport(reports[0].ID); err != nil {
		t.Error(err)
	}

	if all, err := GetAllReports(0); err != nil {
		t.Error(err)
	} else if len(all) != 2 {
		t.Error("Result is wrong.")
	}
}
