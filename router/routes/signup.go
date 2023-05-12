package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-playground/database"
	"net/http"
)

func SignupRoute(r *mux.Router, db database.DBImplementation) error {
	r.HandleFunc("/signup", signupPostHandler(db)).Methods("POST")

	return nil
}

func signupPostHandler(db database.DBImplementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials database.Credentials

		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			return
		}
		token, err := db.SignupUser(credentials)
		if err != nil {
			panic(err)
		}

		if err := json.NewEncoder(w).Encode(token); err != nil {
			panic(err)
		}
	}
}
