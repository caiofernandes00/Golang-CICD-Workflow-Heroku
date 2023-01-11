package db

import (
	"database/sql"
	"fmt"
	"golang-cicd-workflow-heroku/src/util"
	"log"

	// postgres driver
	_ "github.com/lib/pq"
)

func Connect(config util.Config) *sql.DB {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	fmt.Println("Connected to database")
	return db
}
