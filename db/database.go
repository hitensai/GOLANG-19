package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Initdb() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("error while opening connection to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createEventDb()
}

func createEventDb() {
	createEventQuery := `
	CREATE TABLE IF NOT EXISTS events (
	  id INTEGER PRIMARYKEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  datetime DATETIME NOT NULL,
	  user_id INTEGER
	)`

	_, err := DB.Exec(createEventQuery)

	if err != nil {
		panic("error while creating table in database")
	}
}