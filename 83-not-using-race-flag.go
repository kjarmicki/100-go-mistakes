package main

import "fmt"

/*
 * Go is able to detect some data races when a program is run with -race flag.
 * It instruments the code and creates memory and execution time overhead, so it should only be used in non-production scenarios.
 * It cannot catch a false positive - if it detects a data racem there *is* a data race.
 * It can sometimes lead to false negatives (miss actual data race) - this can be alleviated by looping the suspected race part.
 */

func DataRaceFlagSimple() {
	i := 0
	go func() {
		i++
	}()
	fmt.Println(i)
}
