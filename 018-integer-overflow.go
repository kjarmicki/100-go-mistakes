package main

import (
	"fmt"
	"math"
)

func IntegerOverflow() {
	overflowed := math.MaxInt64
	overflowed++
	fmt.Println(overflowed) // -9223372036854775808. no panic, nothing.
}
