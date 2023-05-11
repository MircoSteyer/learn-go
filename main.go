package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go-playground/database"
	"go-playground/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := database.Start()
	if err != nil {
		panic(err)
	}

	router.Start(db)
}
