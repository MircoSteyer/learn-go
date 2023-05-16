package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
	"testing"
	"time"
)

var mockUsers = []User{
	{
		ID:        1,
		CreatedAt: time.Now(),
		Credentials: Credentials{
			Username: "testuser1",
			Password: "testpassword1",
		},
	},
	{
		ID:        2,
		CreatedAt: time.Now(),
		Credentials: Credentials{
			Username: "testuser2",
			Password: "testpassword2",
		},
	},
}

func setupTests(users []User) (Database, *sqlmock.Rows, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at"})
	for _, user := range users {
		rows.AddRow(user.ID, user.Username, user.Password, user.CreatedAt)
	}

	return Database{db}, rows, mock, err
}

func TestDatabase_GetManyUsers(t *testing.T) {
	users := mockUsers
	db, rows, mock, err := setupTests(users)
	if err != nil {
		t.Errorf("an error '%s' was not expected when setting up the test", err)
	}

	// they can be regular expressions if the SQL needs to be less specific
	// e.g. mock.ExpectQuery("^SELECT (.+) FROM users$").WillReturnRows(rows)
	mock.ExpectQuery("SELECT id, username, password, created_at FROM users").WillReturnRows(rows)

	actualUsers, err := db.GetManyUsers()
	if err != nil {
		t.Errorf("an error '%s' was not expected when getting many users", err)
	}

	if !reflect.DeepEqual(users, mockUsers) {
		t.Errorf("expected %v, got %v", users, actualUsers)
	}

	mock.ExpectQuery("SELECT id, username, password, created_at FROM users").WillReturnError(sql.ErrNoRows)

	_, err = db.GetManyUsers()
	if err == nil {
		t.Errorf("an error was expected when executing on an empty table")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDatabase_GetSingleUser(t *testing.T) {
	var users []User
	userId := 2

	for _, mockUser := range mockUsers {
		if mockUser.ID == userId {
			users = append(users, mockUser)
		}
	}

	db, rows, mock, err := setupTests(users)
	if err != nil {
		t.Errorf("an error '%s' was not expected when setting up the test", err)
	}

	mock.ExpectQuery("SELECT id, username, password, created_at FROM users WHERE id = ?").
		WithArgs(userId).
		WillReturnRows(rows)

	actualUser, err := db.GetSingleUser(userId)
	if err != nil {
		t.Errorf("an error '%s' was not expected when getting many users", err)
	}

	if !reflect.DeepEqual(actualUser, users[0]) {
		t.Errorf("expected %v, got %v", users[0], actualUser)
	}

	mock.ExpectQuery("SELECT id, username, password, created_at FROM users WHERE id = ?").WillReturnError(sql.ErrNoRows)

	_, err = db.GetSingleUser(userId)
	if err == nil {
		t.Errorf("an error was expected when executing on an empty table")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
