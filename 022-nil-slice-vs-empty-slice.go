package main

import (
	"encoding/json"
	"fmt"
)

func NilSliceVsEmptySlice() {
	log := func(i int, s []string) {
		fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
	}

	var s []string
	log(1, s)          // empty=true nil=true
	s = append(s, "a") // works, even though it's a nil slice

	s = []string(nil)
	log(2, s) // empty=true nil=true

	s = []string{}
	log(3, s) // empty=false nil=false

	s = make([]string, 0)
	log(4, s) // empty=false nil=false

	// nil slice is marshalled differently than empty slice
	var a []string
	b := []string{}
	jsn, _ := json.Marshal(struct {
		A []string
		B []string
	}{
		A: a,
		B: b,
	})
	fmt.Println(string(jsn)) // {A: null, B: []}

	// rule of thumb:
	// use []string if output size is uncertain and can be zero
	// use make([]string, length) if output size is certain
}
