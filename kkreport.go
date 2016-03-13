package kkreport

import (
	"time"

	"github.com/jackc/pgx"
	"github.com/satori/go.uuid"
)

// Report model.
type Report struct {
	ID     string `json:"id"`
	UserID int32  `json:"userid"`
	At     int32  `json:"at"`
	Reason int16  `json:"reason"`
	Handle bool   `json:"handle"`
	Value  string `json:"value"`
}

// Use the pool to do further operations.
func Use(pool *pgx.ConnPool) error {
	dbPool = pool
	return prepareDB()
}

// InsertReport to insert a report.
func InsertReport(userid int32, reason int16, value string) error {
	var one Report
	one.ID = uuid.NewV1().String()
	one.At = int32(time.Now().Unix())
	one.Reason = reason
	one.Value = value
	one.UserID = userid

	return insertReport(&one)
}

// HandleReport to mark a report as handled.
func HandleReport(id string) error {
	return handleOne(id)
}

// DeleteReport to delete a report.
func DeleteReport(id string) error {
	return deleteOne(id)
}

// GetAllReports to get all reports.
// utime the unixtime, the repots will be got after that time.
func GetAllReports(utime int32) ([]Report, error) {
	return getAll(utime)
}

// GetUnhandledReports to get the unhandled reports.
// utime the unixtime, the repots will be got after that time.
func GetUnhandledReports(utime int32) ([]Report, error) {
	return getUnHandled(utime)
}

// GetHandledReports to get handled reports.
// utime the unixtime, the repots will be got after that time.
func GetHandledReports(utime int32) ([]Report, error) {
	return getHandled(utime)
}
