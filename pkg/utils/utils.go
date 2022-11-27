package utils

import (
	"encoding/json"
	"go-crud-MongoDB/pkg/env"
	"go-crud-MongoDB/pkg/models"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		env.ErrorLog.Println(err)
		return err
	}

	if err := json.Unmarshal(body, x); err != nil {
		env.ErrorLog.Println(err)
		return err
	}
	return nil
}

func ParseUser(r *http.Request, user *models.User) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		env.ErrorLog.Println(err)
		return err
	}

	if err := json.Unmarshal(body, user); err != nil {
		env.ErrorLog.Println(err)
		return err
	}

	return nil
}

func ServeError(w http.ResponseWriter, err error, status int, message string) {
	env.ErrorLog.Println(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
