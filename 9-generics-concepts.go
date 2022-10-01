package main

/*
 * Constraints
 */

// map keys must be comparable (usable with == and != operators), hence the `comparable` constraint
func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}

// constraints are interfaces that can contain arbitrary methods or types
type customComparable interface {
	~int | ~string // ~T means any type with underlying type T
}

type underlyingInt int

func getKeysAlt[K customComparable, V any](m map[K]V) []K {
	// same implementation as getKeys
	return nil
}

// structs can be parametrized too
type Node[T any] struct {
	Val  T
	next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

func GenericConcepts() {
	getKeys(map[int]string{1: "1"})
	getKeys(map[string]int{"1": 1})
	getKeysAlt(map[underlyingInt]int{1: 1})
}

/*
 * Generics use cases:
 * - data structures -> lists, trees, heaps don't need to know about the type they're holding and work well with generics
 * - functions operating on slices, maps and channels -> e.g. function that merges two channels into one
 */

 