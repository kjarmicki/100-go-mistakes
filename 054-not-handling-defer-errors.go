package main

import (
	"database/sql"
	"log"
)

/*
 * Don't ignore errors that happen in deferred functions.
 * Like regular errors, they can be handled or propagated to calling function.
 */

func getBalanceNoHandling(db *sql.DB, clientId string) (float32, error) {
	rows, err := db.Query("...", clientId)
	if err != nil {
		return 0, err
	}
	defer rows.Close() // not optimal - rows.Close() returns an error that we're ignoring here
	return 0, nil
}

func getBalanceLogHandling(db *sql.DB, clientId string) (float32, error) {
	rows, err := db.Query("...", clientId)
	if err != nil {
		return 0, err
	}
	// better - handle an error in the deferred function
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()
	return 0, nil
}

func getBalancePropagate(db *sql.DB, clientId string) (count float32, err error) {
	rows, err := db.Query("...", clientId)
	if err != nil {
		return 0, err
	}
	// alternative - capture err variable and assign it before returning
	defer func() {
		err = rows.Close()
	}()
	return 0, nil
}
