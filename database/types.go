package database

import (
	"database/sql"
	"time"
)

type Database struct {
	*sql.DB
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type DBImplementation interface {
	GetManyUsers() ([]User, error)
	GetSingleUser() (User, error)
}
