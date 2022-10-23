package main

import (
	"fmt"
	"strings"
)

func InefficientStringConcat() {
	values := []string{"a", "large", "slice", "of", "strings"}
	concat := ""
	for _, str := range values {
		concat += str // inefficient, will re-allocate the entire string each time because strings are immutable
	}

	sb := strings.Builder{}
	sb.Grow(len(values))
	for _, str := range values {
		sb.WriteString(str) // appends to an internal buffer which minimizes memory re-allocation
	}
	fmt.Println(sb.String())
}
