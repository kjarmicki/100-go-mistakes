package main

import (
	"os"
	"testing"
)

// utility functions

type TestCustomer struct {
	name string
}

// in testing, there's no point in returning errors from utility functions.
// instead, fail the test right there.
func createTestCustomer(t *testing.T, name string) *TestCustomer {
	if name == "" {
		t.Fatal("test customer name is required")
	}
	return &TestCustomer{
		name: name,
	}
}

// setup and teardown

// one-shot: use simple function call for before and deferred call for after
func TestAnything(t *testing.T) {
	beforeTest()
	defer afterTest()
	// do the test
}

func beforeTest() {

}

func afterTest() {

}

// or like that:
func TestAnythingAgain(t *testing.T) {
	beforeAndAfter(t)
	// do the test
}

func beforeAndAfter(t *testing.T) {
	beforeTest()
	t.Cleanup(func() {
		afterTest()
	})
}

// setup and teardown per package: use TestMain
func TestMain(m *testing.M) {
	beforeTest()
	code := m.Run()
	afterTest()
	os.Exit(code)
}
