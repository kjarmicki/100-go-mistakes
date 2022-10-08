package main

import "fmt"

func SliceCopies() {
	src := []int{1, 2, 3}
	var dst []int
	copy(dst, src)
	fmt.Println(dst) // surprise! it's []

	// result of copying is always a min of src and dst lengths
	dst2 := make([]int, len(src))
	copy(dst2, src)
	fmt.Println(dst2)
}
