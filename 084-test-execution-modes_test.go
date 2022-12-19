package main

import (
	"fmt"
	"testing"
)

/*
 * Go tests can be run with -parallel flag.
 * With that, every test marked with t.Parallel() will be run in parallel.
 * Tests not marked with t.Parallel() will be run before those that are.
 */

func TestInParallelFirst(t *testing.T) {
	t.Parallel()
	fmt.Println("this test will run in parallel")
}

func TestInParallelSecond(t *testing.T) {
	t.Parallel()
	fmt.Println("this test will also run in parallel")
}

/*
 * There's also a -shuffle=[seed] flag.
 * When it's turned on, the execution order of tests will be randomized.
 * Randomization is done using seed to be able to reproduct the results.
 * For example, if go test -shuffle=2435762345 (random number) fails,
 * it can be re-run locally using the same seed.
 */
