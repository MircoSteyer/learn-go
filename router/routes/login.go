package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-playground/database"
	"go-playground/router/middleware"
	"net/http"
)

func LoginRoute(r *mux.Router, db database.DBImplementation) error {
	r.HandleFunc("/login", middleware.Chain(loginPostHandler(db), middleware.Logging())).Methods("POST")

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
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		if err := json.NewEncoder(w).Encode(token); err != nil {
			panic(err)
		}
	}
}
