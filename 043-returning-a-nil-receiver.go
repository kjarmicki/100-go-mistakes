package main

import (
	"errors"
	"fmt"
	"strings"
)

type NilReceiver struct{}

func (n *NilReceiver) Bar() string {
	return "bar"
}

/*
 * In Go, method is just a syntactic sugar for a function with first parameter being a receiver.
 * So even if receiver is nil, a method call still compiles and runs if it doesn't reference the receiver inside.
 */

func UsingANilReceiver() {
	var n *NilReceiver
	if n == nil { // true
		fmt.Println(n.Bar()) // works, prints "bar"
	}
}

type MultiError struct {
	errors []error
}

func (m *MultiError) Add(err error) {
	m.errors = append(m.errors, err)
}

func (m *MultiError) Error() string {
	messages := make([]string, len(m.errors))
	for i, err := range m.errors {
		messages[i] = err.Error()
	}
	return strings.Join(messages, "; ")
}

func validateCustomer() error {
	var m *MultiError
	// do validations
	if false {
		m.Add(errors.New("this will not happen"))
	}
	return m
}

/*
 * In Go, an interface is a dispatch wrapper.
 * Here, the wrapped struct is nil (the MultiError pointer) but the wrappper is not (the error interface).
 * [[*MultiError] error]
 *   ^ nil        ^not nil
 * Hence, the err != nil returns true (because wrapper is not nil) but print outputs nil (because wrapped pointer is nil).
 */

func NilPointerCheck() {
	err := validateCustomer()
	if err != nil { // true
		fmt.Printf("customer is invalid: %v", err) // err is <nil>
	}
}
