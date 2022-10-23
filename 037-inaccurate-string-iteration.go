package main

import (
	"fmt"
	"unicode/utf8"
)

func InaccurateStringIteration() {
	s := "hęllo"
	for i, r := range s {
		fmt.Printf("byte at position %d: %c\n", i, s[i]) // prints the byte
		fmt.Printf("rune at position %d: %c\n", i, r)    // prints the rune
	}
	fmt.Printf("len=%d\n", len(s))         // 6 instead of 5 due to len() counting amount of byts, not length of the string
	fmt.Println(utf8.RuneCountInString(s)) // 5

	// to access a single rune we can't simply use an index
	fmt.Println(string(s[4]))         // wrong, "l"
	fmt.Println(string([]rune(s)[4])) // ok, "o" - a slice of runes can be accessed through an index
}
