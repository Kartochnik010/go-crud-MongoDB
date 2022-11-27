package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"go-crud-MongoDB/pkg/models"
	"go-crud-MongoDB/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	type UserController struct {
//	    BaseController
//	}
func GetAllUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllUsers()
		utils.ServeError(w, err, http.StatusBadRequest, "Couldn't get users")

		json.NewEncoder(w).Encode(users)
	}
}

func GetUserByLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := strings.Split(r.URL.Path, "/")[1]

		user, err := models.GetUserByLogin(query)
		if err != nil {
			utils.ServeError(w, err, http.StatusBadRequest, "Couldn't get user")
		}

		json.NewEncoder(w).Encode(user)
	}
}

func CreateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := utils.ParseUser(r, &user)
		if err != nil {
			utils.ServeError(w, err, http.StatusBadRequest, "Error parsing user")
			return
		}
		// validate user???
		user.TimeStamp = primitive.NewObjectID().Timestamp()
		user.ID = primitive.NewObjectID()

		err = models.CreateUser(user)
		if err != nil {
			utils.ServeError(w, err, http.StatusBadRequest, "Error creating user")
			return
		}

		res, err := json.Marshal(user)
		if err != nil {
			utils.ServeError(w, err, http.StatusBadRequest, "Error marshaling user")
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func UpdateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
