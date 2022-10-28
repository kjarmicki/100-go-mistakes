package main

import (
	"errors"
	"fmt"
	"net/http"
)

/*
 * In Go, errors can be wrapped. Wrapping serves two main purposes:
 * - adding context to an error
 * - marking an error with specific type
 *
 * errors.Is() and errors.As() are useful for digging into wrapped errors to check if they contain an error of specific type
 */

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransationAmount(transactionId string) (float32, error) {
	if len(transactionId) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionId)
	}

	amount, err := getTransactionAmountFromDb(transactionId)
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction %s: %w", transactionId, err)
	}
	return amount, nil
}

func getTransactionAmountFromDb(transactionId string) (float32, error) {
	return 0.0, transientError{err: errors.New("the db is down")}
}

func handler(w http.ResponseWriter, r *http.Request) {
	transactionId := r.URL.Query().Get("transaction")

	amount, err := getTransationAmount(transactionId)
	if err != nil {
		if errors.As(err, &transientError{}) { // this function checks if the error is or wraps transientError at any level
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_ = amount
}
