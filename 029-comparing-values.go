package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	id string
}

type customerF struct {
	id         string
	operations []float64
}

func ComparingValues() {
	cust1 := customer{id: "x"}
	cust2 := customer{id: "x"}
	fmt.Println(cust1 == cust2) // true - two structs can be compared and it will happen by value

	custF1 := customerF{id: "x", operations: []float64{1.}}
	custF2 := customerF{id: "x", operations: []float64{1.}}
	// fmt.Println(custF1 == custF2)
	// ^ this won't even compile because == and != operators don't work for slices or maps
	// things that can be compared: booleans, numbers, strings, channels, interfaces, pointers, structs, arrays
	// things that can't be compared: slices, maps (and also structs that contain them)
	fmt.Println(reflect.DeepEqual(custF1, custF2)) // this works fine, but it's way slower
}
