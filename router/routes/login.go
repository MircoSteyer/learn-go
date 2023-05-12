package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-playground/database"
	"net/http"
)

func LoginRoute(r *mux.Router, db database.DBImplementation) error {
	r.HandleFunc("/login", loginPostHandler(db)).Methods("POST")

	return nil
}

func loginPostHandler(db database.DBImplementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials database.Credentials

		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			return
		}
		token, err := db.LoginUser(credentials)
		if err != nil {
			panic(err)
		}

		if err := json.NewEncoder(w).Encode(token); err != nil {
			panic(err)
		}
	}
}
