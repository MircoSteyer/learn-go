package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-playground/database"
	"net/http"
	"strconv"
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

		if err := json.NewEncoder(w).Encode(users); err != nil {
			panic(err)
		}
	}
}

func singleUserGetHandler(db database.DBImplementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringId := mux.Vars(r)["id"]

		id, err := strconv.Atoi(stringId)
		if err != nil {
			panic(err)
		}

		user, err := db.GetSingleUser(id)
		if err != nil {
			panic(err)
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
	}
}
