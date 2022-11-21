package main

import (
	"context"
	"database/sql"
	"log"
)

/*
 * When iterating over SQL rows it's easy to miss certain type of an error.
 */

func readEmployeeNames(ctx context.Context, db *sql.DB) ([]string, error) {
	rows, err := db.QueryContext(ctx, "SELECT name FROM employees")
	if err != nil {
		return []string{}, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	names := make([]string, 0)
	// beware: the loop may break when there are no more rows
	// OR when there is an error while preparing the next row
	// therefore it's a must to check rows.Err()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return []string{}, err
		}
		names = append(names, name)
	}

	// this check is easy to miss but without it's possible to return incomplete amount of rows
	if err := rows.Err(); err != nil {
		return []string{}, err
	}

	return names, nil
}
