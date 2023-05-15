package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-playground/database"
	"go-playground/router/middleware"
	"net/http"
	"strconv"
)

func UsersRoute(r *mux.Router, db database.DBImplementation) error {
	r.HandleFunc("/users", middleware.Chain(
		manyUsersGetHandler(db),
		middleware.Auth(),
		middleware.Logging())).
		Methods("GET")

	r.HandleFunc("/users/me", middleware.Chain(
		singleUserGetHandler(db),
		middleware.Auth(),
		middleware.Logging())).
		Methods("GET")

	r.HandleFunc("/users/{id}", middleware.Chain(
		singleUserGetHandler(db),
		middleware.Auth(),
		middleware.Logging())).
		Methods("GET")

	return nil
}

func manyUsersGetHandler(db database.DBImplementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetManyUsers()
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
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
			http.Error(w, "Invalid ID: identifier has to be of type int", http.StatusBadRequest)
			return
		}

		user, err := db.GetSingleUser(id)
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
	}
}
