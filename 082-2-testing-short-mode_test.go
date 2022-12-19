package main

import (
	"fmt"
	"testing"
)

/*
 * While running tests, it's possible to target only those that are short-running.
 * go test -short will set the testing.Short() return value to true
 * which can be used in the test body.
 */

func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test")
		return
	}
	fmt.Println("perform a long-running test")
}
