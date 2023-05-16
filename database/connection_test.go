package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestStart(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	dbConnectorMock := func(_ string, _ string) (*sql.DB, error) {
		return db, nil
	}
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectPing()

	_, err = Start(dbConnectorMock)

	if err != nil {
		t.Errorf("an error '%s' was not expected when starting the database", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
