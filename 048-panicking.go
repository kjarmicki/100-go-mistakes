package main

import "regexp"

/*
 * In Go, panicking is appropriate under two circumstances:
 * - an obvious programmer error
 * - a failure to initialize a dependency
 */

func anObviousProgrammerError(httpStatusCode int) {
	if httpStatusCode < 100 || httpStatusCode > 999 {
		panic("invalid status code")
	}
}

func aFailureToInitializeADependency() *regexp.Regexp {
	return regexp.MustCompile("jibberish")
}
