package main

import "database/sql"

/*
 * Simply connecting to an SQL database may be insufficient to conclude that the database is ready to respond.
 */

func EnsureSqlConnection() {
	// for some drivers, sql.Open doesn't establish a connection, it only prepares db for later use (like with db.Query)
	db, err := sql.Open("mysql", "dsn")
	if err != nil {
		panic(err)
	}

	// to be really sure that the database is ready to take requests, use the Ping method
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
