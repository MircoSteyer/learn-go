package database

import (
	"fmt"
	"os"
)

func Start(connector DBConnector) (DBImplementation, error) {
	var (
		user     = os.Getenv("MYSQL_USER")
		password = os.Getenv("MYSQL_PASSWORD")
		address  = os.Getenv("MYSQL_ADDRESS")
		dbName   = os.Getenv("MYSQL_DATABASE")
	)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, address, dbName)

	db, err := connector("mysql", connectionString)
	database := &Database{DB: db}
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
