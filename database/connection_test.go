package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestStart(t *testing.T) {
	dbMock, expectationMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	dbConnectorMock := func(_ string, _ string) (*sql.DB, error) {
		return dbMock, nil
	}
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expectationMock.ExpectPing()

	_, err = Start(dbConnectorMock)

	if err != nil {
		t.Errorf("an error '%s' was not expected when starting the database", err)
	}

	if err := expectationMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
