package main

import "fmt"

func RangeLoopNotInfinite() {
	s := []int{0, 1, 2}
	for range s {
		s = append(s, 10)
	}
	fmt.Println("range loop finished, s len:", len(s))
	// the loop won't be infinite because in range loops expression is evaluated only once, at the start of the loop

	// it's different for the classic for loop though: it evaluates slice length every time
	for i := 0; i < len(s); i++ {
		s = append(s, 10)
		if i == 1000 {
			break
		}
	}
	fmt.Println("classic loop finished, s len:", len(s))
}

func RangeLoopNoModification() {
	a := [3]int{0, 1, 2}  // notice: array, not a slice
	for i, v := range a { // entire array is copied for iteration
		a[2] = 10
		if i == 2 {
			fmt.Println(v) // so this is 2, not 10
		}
	}
}
