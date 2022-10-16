package main

func InefficientMapInit() {
	// it's more efficient to specify initial size of a map, if known, than to let it grow
	// because growing and moving elements between buckets is expensive

	_ = make(map[string]int, 1_000_000)
}
