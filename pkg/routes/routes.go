package routes

import (
	"go-crud-MongoDB/pkg/controllers"

	"github.com/gorilla/mux"
)

var (
	RegisterRoutes = func(r *mux.Router) {
		r.HandleFunc("/", controllers.GetAllUsers()).Methods("GET")           // get all users
		r.HandleFunc("/", controllers.CreateUser()).Methods("POST")           // create a user
		r.HandleFunc("/{login}", controllers.GetUserByLogin()).Methods("GET") // get user by id
		r.HandleFunc("/{login}", controllers.UpdateUser()).Methods("PUT")     // update a user by id
		r.HandleFunc("/{login}", controllers.DeleteUser()).Methods("DELETE")  // delete a user by id
	}
)
