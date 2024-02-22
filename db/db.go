package db

import (
	"database/sql" //this is an interface that is implemented by various db drivers . this makes our code db agnostic

	_ "github.com/glebarez/go-sqlite" // get this
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	// createEventsTable := `
	// CREATE TABLE IF NOT EXISTS events (
	//     id INTEGER PRIMARY KEY AUTOINCREMENT,
	//     name TEXT NOT NULL,
	//     description TEXT NOT NULL,
	//     location TEXT NOT NULL,
	//     dateTime DATETIME NOT NULL,
	//     user_id INTEGER
	// )
	// `

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
