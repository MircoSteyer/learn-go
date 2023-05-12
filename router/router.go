package router

import (
	"github.com/gorilla/mux"
	"go-playground/database"
	"go-playground/router/routes"
	"net/http"
)

func Start(db database.DBImplementation) {
	r := mux.NewRouter()
	err := routes.UsersRoute(r, db)
	err = routes.SignupRoute(r, db)
	err = routes.LoginRoute(r, db)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
