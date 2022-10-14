package main

import (
	"fmt"
	"runtime"
)

type LeakyFoo struct {
	v []byte
}

type Keeper func([]LeakyFoo) []LeakyFoo

func LeakyFooCheck() {
	performLeakCheck(copyKeepFirstTwoElementsOnly)
	fmt.Println("------")
	performLeakCheck(leakyKeepFirstTwoElementsOnly)
}

func performLeakCheck(keeper Keeper) {
	foos := make([]LeakyFoo, 1_000)
	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = LeakyFoo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	two := keeper(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func leakyKeepFirstTwoElementsOnly(foos []LeakyFoo) []LeakyFoo {
	// this will leak (1000MB istead of 2MB allocated memory)
	// because backing array is still referenced by the returned slice
	return foos[:2]
}

func copyKeepFirstTwoElementsOnly(foos []LeakyFoo) []LeakyFoo {
	// this is not going to leak because there's no backing array anymore
	output := make([]LeakyFoo, 2)
	copy(output, foos)
	return output
}
