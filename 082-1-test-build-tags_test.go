//go:build integration

/*
 * Separating tests is one of the primary use cases for build tags.
 * With this build tag:
 * - go test -v . will run only the test files without build tags
 * - go test --tags=integration -v . will run the test files with the tag AND the ones without any tag
 *
 * In order to run only integration test one must add //go:build !integration in other test files.
 * Serious drawback: absence of any signal that a test has been ignored. Output shows only tests that were executed.
 */

package main

import (
	"fmt"
	"testing"
)

func TestIntegration(t *testing.T) {
	fmt.Println("perform integration tests")
}
