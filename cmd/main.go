package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database, _ := sql.Open("sqlite3", "./nraboy.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
}
