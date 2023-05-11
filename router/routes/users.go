package routes

import (
	"github.com/gorilla/mux"
	"go-playground/database"
	"log"
	"net/http"
)

func UsersRoute(r *mux.Router, db database.DBImplementation) error {
	r.HandleFunc("/users", manyUsersGetHandler(db)).Methods("GET")
	r.HandleFunc("/users/{id}", singleUserGetHandler(db)).Methods("GET")

	return nil
}

func manyUsersGetHandler(db database.DBImplementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetManyUsers()
		if err != nil {
			panic(err)
		}
		log.Println("USERS", users)
	}
}

func singleUserGetHandler(db database.DBImplementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetManyUsers()
		if err != nil {
			panic(err)
		}
		log.Println("USERS", users)
	}
}
