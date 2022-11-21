package main

import "database/sql"

/*
 * Preprared statement is a precompiled query that can be reused.
 * Benefits:
 * - efficiency -> the same statements don't need to be compiled (parsed + optimized + translated) over and over
 * - security -> reduced risk of SQL injection
 */

type SomeRepository struct {
	db       *sql.DB
	findStmt *sql.Stmt
}

func NewSomeRepository(db *sql.DB) (*SomeRepository, error) {
	// prepare and store the statement for later use
	findStmt, err := db.Prepare("SELECT * FROM Some WHERE ID = ?")
	if err != nil {
		return nil, err
	}
	return &SomeRepository{
		db:       db,
		findStmt: findStmt,
	}, nil
}

func (sr *SomeRepository) findOrder(id int) error {
	// use the statement
	rows, err := sr.findStmt.Query(id)
	if err != nil {
		return err
	}
	// proceed with rows
	_ = rows
	return nil
}
