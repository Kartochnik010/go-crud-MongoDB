package env

import (
	"log"
	"os"

	"go-crud-MongoDB/pkg/db"
)

// var App = &models.Application{
// 	InfoLog:  log.New(os.Stdout, "", log.Ltime),
// 	ErrorLog: log.New(os.Stderr, "", log.Lshortfile|log.Ltime),
// 	DB:       db.GetClient(),
// }

var (
	InfoLog  = log.New(os.Stdout, "", log.Ltime)
	ErrorLog = log.New(os.Stderr, "", log.Lshortfile|log.Ltime)
	DB       = db.MustGetClient()
)
