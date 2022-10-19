package main

import "fmt"

func MapInsertDoingIteration() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for i := 0; i < 10; i++ {
		m2 := copyMap(m)
		for k, v := range m2 {
			if v {
				m2[10+k] = true
				/*
				* Inserting an element to a map while iterating over it may result in unpredictable results.
				* A newly added element may be iterated over in the loop afterwards or not.
				* There is no way to force Go to iterate over it or skip it.
				*
				* If there is a need to modify a map during iteration, it can be done with a copy - iterate over original array, insert things into the copy.
				 */
			}
		}
		fmt.Println(m2)
	}
}

func copyMap[K comparable, V any](m map[K]V) map[K]V {
	var copy = make(map[K]V, len(m))
	for k, v := range m {
		copy[k] = v
	}
	return copy
}
