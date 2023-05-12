package database

import (
	"database/sql"
	"time"
)

type Database struct {
	*sql.DB
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID int `json:"id"`
	Credentials
	CreatedAt time.Time `json:"created_at"`
}

type DBUsers interface {
	GetManyUsers() ([]User, error)
	GetSingleUser(id int) (User, error)
}

type Signup interface {
	SignupUser(credentials Credentials) (string, error)
}

type Login interface {
	LoginUser(credentials Credentials) (string, error)
}

type DBImplementation interface {
	DBUsers
	Signup
	Login
}
