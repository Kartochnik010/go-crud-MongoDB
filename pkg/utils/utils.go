package utils

import (
	"encoding/json"
	"go-crud-MongoDB/pkg/env"
	"go-crud-MongoDB/pkg/models"
	"io/ioutil"
	"net/http"
)

type JSON map[string]interface{}

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

func ServeError(w http.ResponseWriter, err error, status int, msg string) {
	env.ErrorLog.Println(err)
	R(w, status, JSON{
		"msg": msg,
		"err": err.Error(),
	})
}

// Writer with an R
// Automatically sets Content-Type to "application/json"
func R(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	e.SetIndent("", "    ")
	e.Encode(response)
}

// Playing around
// Automatically sets Content-Type to "application/json"
type ErrorR struct {
	W        http.ResponseWriter
	Status   int
	Response JSON
}

func (c ErrorR) Run() {
	c.W.Header().Set("Content-Type", "application/json")
	c.W.WriteHeader(c.Status)
	e := json.NewEncoder(c.W)
	e.SetIndent("", "    ")
	e.Encode(c.Response)
}
