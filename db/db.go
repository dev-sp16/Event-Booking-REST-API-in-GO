package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // '_' => We will not use this package directly  
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open( "sqlite3", "api.db" )

	if err != nil {
		panic( "ERROR: Could not connect to the database." )
	}

	DB.SetMaxOpenConns( 10 ) // 10 simultaneous connections
	DB.SetMaxIdleConns( 5 ) // 5 max IDLE conections

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		PASSWORD TEXT NOT NULL
	)
	`
	_, err := DB.Exec( createUsersTable )

	if err != nil {
		panic( err )
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY( user_id ) REFERENCES users(id)
	)
	`
	_, err = DB.Exec( createEventsTable )

	if err != nil {
		panic( err)
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		even_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY( event_id ) REFERENCES events( id ),
		FOREIGN KEY( user_id ) REFERENCES users( id )
	)
	`
	_, err = DB.Exec( createRegistrationsTable )

	if err != nil {
		panic( err)
	}
}