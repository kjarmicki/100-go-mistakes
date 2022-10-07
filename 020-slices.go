package main

import "fmt"

func Slices() {
	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	// slices s1 and s2 reference the same backing array but with different lengths and capacities

	fmt.Println("s1", s1) // [0, 0, 0]
	fmt.Println("s2", s2) // [0, 0]

	s1[1] = 1
	fmt.Println("s1", s1) // [0, 1, 0]
	fmt.Println("s2", s2) // [1, 0]

	s2 = append(s2, 2)
	fmt.Println("s1", s1) // [0, 1, 0]
	fmt.Println("s2", s2) // [1, 0, 2]

	s1 = append(s1, 3)
	fmt.Println("s1", s1) // [0, 1, 0, 3]
	fmt.Println("s2", s2) // [1, 0, 3]
}
