package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go-crud-MongoDB/pkg/env"
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
		if err != nil {
			utils.ServeError(w, err, http.StatusBadRequest, "Couldn't get users")
			return
		}
		utils.R(w, 200, users)
	}
}

func GetUserByLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// looking for better alternative getting query from path
		query := strings.Split(r.URL.Path, "/")[1]

		user, err := models.GetUserByLogin(query)
		if err != nil {
			utils.ServeError(w, err, http.StatusBadRequest, "Couldn't get user")
			return
		}

		utils.R(w, 200, user)
	}
}

func CreateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		// parsing user from request body
		err := utils.ParseUser(r, &user)
		if err != nil {
			utils.ServeError(w, err, http.StatusBadRequest, "Error parsing user")
			return
		}

		// checking for unique username
		err = user.Validate()
		if err != nil {
			utils.ServeError(w, err, http.StatusNotAcceptable, "Error validating user")
			return
		}

		// saving user, but we don't know the _id yet
		InsertOneResult, err := models.CreateUser(user)
		if err != nil {
			utils.ServeError(w, err, http.StatusBadRequest, "Error creating user")
			return
		}

		// Getting user's generated _id as a response from MongoDB and filling the missing fields of struct.
		// Filling missing fields in the struct seems unreasonable, but we are supposed to see what is put to the database
		if user.ID.IsZero() {
			user.ID, err = primitive.ObjectIDFromHex(fmt.Sprintf("%s", InsertOneResult.InsertedID)[10:34])
			if err != nil {
				fmt.Println(fmt.Sprintf("%s", InsertOneResult.InsertedID)[10:34])
				utils.ServeError(w, err, http.StatusAccepted, "User inserted, but failed to show")
				return
			}
		}
		user.TimeStamp = user.ID.Timestamp()

		// writing response
		utils.R(w, http.StatusCreated, user)
	}
}

func UpdateUserByLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// func  method(...error)  {
// 	for range append {"err" : err.Error()}
// }

func DeleteUserByLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := strings.Split(r.URL.Path, "/")[1]
		RowsDeleted, err := models.DeleteUserByLogin(query)
		if err != nil {
			utils.ServeError(w, err, http.StatusConflict, "Couldn't delete user")
			return
		}

		if RowsDeleted == 0 {
			utils.ServeError(w, fmt.Errorf("0 rows affected"), http.StatusConflict, "Couldn't delete user")
			return
		}

		// another style of sending response
		// weird, but I like it
		utils.ErrorR{
			W:      w,
			Status: http.StatusOK,
			Response: utils.JSON{
				"msg":          "User successfully deleted.",
				"rows_deleted": RowsDeleted,
			},
		}.Run()

	}
}

func DropDatabase() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.DB.Database("novye").Collection("users").Drop(context.TODO())
	}
}
